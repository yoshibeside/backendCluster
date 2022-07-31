package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/frankenbeanies/randhex"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID    int    `json: "id"`
	Title string `json: "title"`
	Done  bool   `json: "done"`
	Body  string `json: "body"`
}

type Node struct {
	ID    float64 `json: "id" bson:"id"`
	X     float64 `json: "x" bson: "x"`
	Y     float64 `json: "y" bson: "y"`
	Color string  `json: "colortype" bson: "colortype"`
}

type KumpulID struct {
	ID1   float64 `json: "id1"`
	ID2   float64 `json: "id2"`
	Jarak float64 `json: "distance"`
	Color string  `json: "colortype"`
}

type Result struct {
	Result []Node `json: "result"`
}

const connectionString = "mongodb+srv://kruskaldb:6kh3gNwtJdwSbxd7@cluster1.q4lutha.mongodb.net/?retryWrites=true&w=majority"

var collection *mongo.Collection

func init() {
	// client options
	clientOption := options.Client().ApplyURI(connectionString)

	// Connectiong to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongodb Connection Success")

	collection = client.Database("Data").Collection("Nothing")

	fmt.Println("Collection instance is ready")
}

func insertOneNodes(result Result) {
	koleksi, err := collection.InsertOne(context.Background(), result)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 node in db with id ", koleksi.InsertedID)
}

func gettingNodes() []Result {

	temp := []Result{}

	finding, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	var results []bson.M
	if err = finding.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	for _, result := range results {
		blabla := &Result{}
		output, err := json.Marshal(result)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(output, &blabla)
		if err != nil {
			panic(err)
		}
		temp = append(temp, *blabla)
		fmt.Println("just 1 next....\n", temp)
	}
	return temp

}

func main() {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	nodes := []Node{}

	app.Get("/bring", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Post("/api/:cluster/data", func(c *fiber.Ctx) error {
		// Menerima input file
		file, err := c.FormFile("file")
		nodes = nodes[:0]
		if err != nil {
			fmt.Println(c)
			return c.SendString("FAILED succesfully")
		}

		c.SaveFile(file, ("../uploaded/nodes.xlsx"))

		data, err := excelize.OpenFile("../uploaded/nodes.xlsx")

		sheetName := data.WorkBook.Sheets.Sheet[0].Name
		// Isi file dimasukkan ke dalam array of node dalam bariabel nodes
		maxRow := 0
		for i, row := range data.GetRows(sheetName) {
			node := &Node{}

			for j, colCell := range row {
				if i >= 1 {
					if j == 0 {
						node.ID, err = strconv.ParseFloat(colCell, 64)
						maxRow = int(node.ID)
					} else if j == 1 {
						node.X, err = strconv.ParseFloat(colCell, 64)
					} else if j == 2 {
						node.Y, err = strconv.ParseFloat(colCell, 64)
					}
				}
				fmt.Print(colCell, "\t")
			}
			nodes = append(nodes, *node)
			fmt.Println()
		}

		fmt.Print("hasil dari nodes adalah ")
		fmt.Println(nodes)

		if err != nil {
			return err
		}

		// Making colors for clustering result
		cluster, err := c.ParamsInt("cluster")

		if err != nil {
			return err
		}

		if cluster > (len(nodes)-1) || cluster < 1 {
			return errors.New("bad input")
		}

		colors := []string{}
		for l := 0; l < cluster; l++ {
			colors = append(colors, randhex.New().String())
		}

		fmt.Println(colors)
		// Processing the Nodes

		// Creating weight matrix for input nodes
		weightMatrix := [10][10]KumpulID{}

		// Inputing inside the weight matrix calculated distance between nodes.
		for i, objects_i := range nodes {
			for j, objects_j := range nodes {
				if i == j {

				} else if i != 0 && j != 0 && i != j {
					weightMatrix[i-1][j-1].Jarak = processing(objects_i.X, objects_i.Y, objects_j.X, objects_j.Y)
					weightMatrix[i-1][j-1].ID1 = objects_i.ID
					weightMatrix[i-1][j-1].ID2 = objects_j.ID
				}
			}
		}

		// Inputing into a normal array to sort the distance between nodes in ascending order.

		sortingNodes := []KumpulID{}

		for i, array1 := range weightMatrix {
			for j, objects := range array1 {
				if (j > i) && (objects.Jarak != 0) {
					sortingNodes = append(sortingNodes, objects)
				}
			}
		}

		// Time for sorting (selection sort)
		var n = len(sortingNodes)
		for i := 0; i < n; i++ {
			var minIdx = i
			for j := i; j < n; j++ {
				if sortingNodes[j].Jarak < sortingNodes[minIdx].Jarak {
					minIdx = j
				}
			}
			sortingNodes[i], sortingNodes[minIdx] = sortingNodes[minIdx], sortingNodes[i]
		}

		results := Kruskal(sortingNodes, maxRow)
		results = clustering(results, colors)
		nodes = settingColors(nodes, results, colors)
		endResult := Result{}
		endResult.Result = nodes
		insertOneNodes(endResult)
		return c.JSON(nodes)
	})

	app.Get("api/get/data", func(c *fiber.Ctx) error {
		got_it := gettingNodes()
		return c.JSON(got_it)
	})
	log.Fatal(app.Listen(":4000"))
}
