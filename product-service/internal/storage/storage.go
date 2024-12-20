package storage

import (
	"context"
	"fmt"
	pb "product-service/genproto/productpb"
	"product-service/internal/storage/models"
	m "product-service/internal/storage/mongodb"
	"product-service/logger"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepo struct {
	mongo *m.Mongo
}

func NewProductRepository(mongosh *m.Mongo) IProductRepository {
	return &ProductRepo{
		mongo: mongosh,
	}
}

const constanta = "allDataFromongo"

func (db *ProductRepo) CreateProduct(req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	logger.Info("CreateProduct: started: ", req.Name)
	// So'nggi mahsulotni topish
	var lastProduct bson.M
	opts := options.FindOne().SetSort(bson.D{{Key: "bag_id", Value: -1}})
	err := db.mongo.Collection.FindOne(
		context.Background(),
		bson.D{},
		opts,
	).Decode(&lastProduct)

	var newBagId int64
	if err == mongo.ErrNoDocuments {
		newBagId = 1
	} else if err != nil {
		logger.Error("CreateProduct: error finding last product - ", err)
		return nil, fmt.Errorf("failed to find last product: %v", err)
	} else {
		lastBagId, ok := lastProduct["bag_id"].(string)
		if !ok {
			newBagId = 1
		} else {
			intNewBagID, err := strconv.ParseInt(lastBagId, 10, 64)
			if err != nil {
				return nil, err
			}

			newBagId = intNewBagID + 1
		}
	}
	createdAt := time.Now().Format("2006-01-02 15:04:05")

	// Mahsulot hujjatini tayyorlash
	productDocument := bson.D{
		{Key: "image_url", Value: req.ImageUrl},
		{Key: "name", Value: req.Name},
		{Key: "unique_number", Value: req.UniqueNumber},
		{Key: "bag_id", Value: strconv.FormatInt(newBagId, 10)},
		{Key: "price", Value: req.Price},
		{Key: "size", Value: req.Size},
		{Key: "colors", Value: req.Colors},
		{Key: "count", Value: req.Count},
		{Key: "created_at", Value: createdAt},
		{Key: "updated_at", Value: time.Now().String()},
		{Key: "deleted_at", Value: 0},
	}

	// Mahsulotni bazaga qo'shish
	result, err := db.mongo.Collection.InsertOne(context.Background(), productDocument)
	if err != nil {
		logger.Error("CreateProduct: error creating product - ", err)
		return nil, fmt.Errorf("failed to insert product: %v", err)
	}

	fmt.Println(productDocument)

	// MongoDB ObjectID-ni olish
	var productID string
	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		logger.Error("CreateProduct: error converting InsertedID to ObjectID")
		return nil, fmt.Errorf("failed to convert product ID")
	}
	productID = objectID.Hex()

	// Javobni tayyorlash
	resp := &pb.CreateProductResponse{
		Product: &pb.Product{
			ImageUrl:     req.ImageUrl,
			Id:           productID,
			Name:         req.Name,
			UniqueNumber: req.UniqueNumber,
			BagId:        strconv.FormatInt(newBagId, 10),
			Price:        req.Price,
			Size:         req.Size,
			Colors:       req.Colors,
			Count:        req.Count,
			Timestamp: []*pb.CUDP{
				{
					CreatedAt: createdAt,
					UpdatedAt: "",
					DeletedAt: 0,
				},
			},
		},
	}

	logger.Info("CreateProduct: product successfully created: ", resp.Product.Id)
	return resp, nil
}

func (db *ProductRepo) GetProductById(req *pb.GetProductByIdRequest) (*pb.GetProductByIdResponse, error) {

	logger.Info("GetProductById: started - ", req.Id)

	resp := pb.GetProductByIdResponse{}

	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		logger.Error("GetProductById: error with product id (primitive.ObjectID)")
		return nil, err
	}

	filter := bson.M{"_id": objectID, "deleted_at": 0}
	singleResult := db.mongo.Collection.FindOne(context.Background(), filter)
	if err := singleResult.Err(); err != nil {
		logger.Error("GetProductById: error finding product by id - ", singleResult.Err())
		return nil, err
	}

	get := models.MongoProduct{}
	if err := singleResult.Decode(&get); err != nil {
		logger.Error("GetProductById: error decoding product from mongodb -  ", err)
		return nil, err
	}

	product := pb.Product{
		ImageUrl:     get.ImageUrl,
		Id:           get.Id.Hex(),
		Name:         get.Name,
		UniqueNumber: get.UniqueNumber,
		BagId:        get.BagID,
		Price:        get.Price,
		Size:         get.Size,
		Colors:       get.Colors,
		Count:        get.Count,
		Timestamp: []*pb.CUDP{
			{
				CreatedAt: get.CreatedAt,
				UpdatedAt: get.UpdatedAt,
				DeletedAt: get.DeletedAt,
			},
		},
	}

	resp.Product = &product
	return &resp, nil
}

func (db *ProductRepo) GetProductByFilter(req *pb.GetProductByFilterRequest) (*pb.GetProductByFilterResponse, error) {

	logger.Info("GetProductByFilter: started: filter = ", req.Search)

	resp := pb.GetProductByFilterResponse{}

	baseFilter := bson.M{"deleted_at": 0}

	var filter bson.M
	if req.Search == constanta {
		filter = baseFilter
		logger.Info("GetProductByFilter: using only baseFilter")
	} else {
		logger.Info("GetProductByFilter: using filter")
		filter = bson.M{
			"$and": []bson.M{
				baseFilter,
				{
					"$or": []bson.M{
						{"name": bson.M{"$regex": req.Search, "$options": "i"}},
						{"unique_number": bson.M{"$regex": req.Search, "$options": "i"}},
						{"bag_id": bson.M{"$regex": req.Search, "$options": "i"}},
						{"size": bson.M{"$regex": req.Search, "$options": "i"}},
						{"count": bson.M{"$regex": req.Search, "$options": "i"}},
						{
							"$expr": bson.M{
								"$regexMatch": bson.M{
									"input":   bson.M{"$toString": "$deadline"},
									"regex":   req.Search,
									"options": "i",
								},
							},
						},
					},
				},
			},
		}
	}

	cursor, err := db.mongo.Collection.Find(context.Background(), filter)
	if err != nil {
		logger.Error("GetProductByFilter: error getting products - ", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		get := models.MongoProduct{}
		if err := cursor.Decode(&get); err != nil {
			logger.Error("GetProductByFilter: error decoding from mongodb -  ", err)
			return nil, err
		}
		product := pb.Product{
			ImageUrl:     get.ImageUrl,
			Id:           get.Id.Hex(),
			Name:         get.Name,
			UniqueNumber: get.UniqueNumber,
			BagId:        get.BagID,
			Price:        get.Price,
			Size:         get.Size,
			Colors:       get.Colors,
			Count:        get.Count,
			Timestamp: []*pb.CUDP{
				{
					CreatedAt: get.CreatedAt,
					UpdatedAt: get.UpdatedAt,
					DeletedAt: get.DeletedAt,
				},
			},
		}
		resp.Products = append(resp.Products, &product)
	}

	logger.Info("GetProductByFilter: product list fetched, size = ", len(resp.Products))
	return &resp, nil
}

func (db *ProductRepo) UpdateStock(req *pb.UpdateStockRequest) (*pb.UpdateStockResponse, error) {

	logger.Info("UpdateStock: started: ", req.ProductId)

	resp := pb.UpdateStockResponse{}

	objectID, err := primitive.ObjectIDFromHex(req.ProductId)
	if err != nil {
		logger.Error("UpdateStock: error with product id (primitive.ObjectID)")
		return nil, err
	}

	filter := bson.M{"_id": objectID, "deleted_at": 0}

	updatedAt := time.Now().Format("2006-01-02 15:04:05")
	product := models.MongoProduct{}
	err = db.mongo.Collection.FindOne(context.Background(), filter).Decode(&product)
	if err != nil {
		logger.Error("UpdateStock: error finding product - ", err)
		return nil, err
	}

	currentCount := product.Colors[req.ProductColor]
	update := bson.M{
		"$set": bson.D{
			{Key: fmt.Sprintf("colors.%s", req.ProductColor), Value: currentCount - 1},
			{Key: "updated_at", Value: updatedAt},
		},
		"$inc": bson.M{
			"count": -1, // Umumiy count dan 1 kamaytirish
		},
	}

	result, err := db.mongo.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error("UpdateStock: error updating product stock - ", err)
		return nil, err
	}
	if result.ModifiedCount == 0 {
		logger.Warn("UpdateStock: no document was updated - product might already have the same stock count")
		return nil, fmt.Errorf("updateStock: no document was updated")
	}

	resp.IsSuccess = true
	return &resp, nil
}

func (db *ProductRepo) UpdateProduct(req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {

	logger.Info("UpdateProduct: started: ", req.Id)

	resp := pb.UpdateProductResponse{}

	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		logger.Error("UpdateProduct: error with product id (primitive.ObjectID)")
		return nil, err
	}

	filter := bson.M{"_id": objectID, "deleted_at": 0}

	updatedAt := time.Now().Format("2006-01-02 15:04:05")
	update := bson.M{
		"$set": bson.D{
			{Key: "image_url", Value: req.ImageUrl},
			{Key: "name", Value: req.Name},
			{Key: "unique_number", Value: req.UniqueNumber},
			{Key: "bag_id", Value: req.BagId},
			{Key: "price", Value: req.Price},
			{Key: "size", Value: req.Size},
			{Key: "colors", Value: req.Colors},
			{Key: "count", Value: req.Count},
			{Key: "updated_at", Value: updatedAt},
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedProduct models.MongoProduct
	err = db.mongo.Collection.FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&updatedProduct)
	if err != nil {
		logger.Error("UpdateProduct: error updating or re-fetching product -  ", err)
		return nil, err
	}

	resp.Product = &pb.Product{
		ImageUrl:     updatedProduct.ImageUrl,
		Id:           updatedProduct.Id.Hex(),
		Name:         updatedProduct.Name,
		UniqueNumber: updatedProduct.UniqueNumber,
		BagId:        updatedProduct.BagID,
		Price:        updatedProduct.Price,
		Size:         updatedProduct.Size,
		Colors:       updatedProduct.Colors,
		Count:        updatedProduct.Count,
		Timestamp: []*pb.CUDP{
			{
				CreatedAt: updatedProduct.CreatedAt,
				UpdatedAt: updatedProduct.UpdatedAt,
				DeletedAt: updatedProduct.DeletedAt,
			},
		},
	}
	logger.Info("UpdateProduct: product successfully updated: ", req.Id)
	return &resp, nil
}

func (db *ProductRepo) DeleteProduct(req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {

	logger.Info("DeleteProduct: started: ", req.Id)

	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		logger.Error("DeleteProduct: error with product id (primitive.ObjectID)")
		return nil, err
	}

	filter := bson.M{"_id": objectID, "deleted_at": 0}
	update := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now().Unix(),
		},
	}

	var deletedProduct models.DeletedProduct
	err = db.mongo.Collection.FindOneAndUpdate(context.Background(), filter, update).Decode(&deletedProduct)
	if err != nil {
		logger.Error("DeleteProduct: error deleting or re-fetching product -  ", err)
		return nil, err
	}

	resp := pb.DeleteProductResponse{
		Name:    deletedProduct.Name,
		Message: "deleted",
	}
	logger.Info("DeleteProduct: product successfully deleted: ", req.Id)
	return &resp, nil
}
func (db *ProductRepo) AddProductCountAndColor(req *pb.AddProductCountAndColorRequest) (*pb.AddProductCountAndColorResponse, error) {
	ctx := context.Background()
	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	// UpdateOne uchun update operatsiyalarini yaratish
	updates := bson.D{}

	// Ranglarni qayta ishlash
	for color, value := range req.Colors {
		// Har bir rang uchun increment qo'shish
		updates = append(updates, primitive.E{
			Key:   "colors." + color,
			Value: value,
		})
	}

	// MongoDB ga yangilash so'rovi
	_, err = db.mongo.Collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$inc": updates},
	)
	if err != nil {
		return nil, err
	}

	return &pb.AddProductCountAndColorResponse{
		Status:  true,
		Message: "success",
	}, nil
}
