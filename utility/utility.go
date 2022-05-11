package utility

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	bolt "go.etcd.io/bbolt"
)

func Init(file string) {
	db, err := bolt.Open(file, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		//handle error
		log.Fatal(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("Video_game"))
		return err
	})
	db.Close()
}

func Open(file string) *bolt.DB {
	db, err := bolt.Open(file, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		//handle error
		log.Fatal(err)
	}
	return db
}

func Set(db *bolt.DB, bucket, key, value string) {
	db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte(bucket))
		err := b.Put([]byte(key), []byte(value))
		return err
	})
}

func Get(db *bolt.DB, bucket, key string) string {
	var s []byte
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		s = b.Get([]byte(key))
		if s == nil {
			s = []byte("0")
		}
		return nil
	})

	return string(s)
}

func Del(db *bolt.DB, bucket, key string) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		b.Delete([]byte(key))
		return nil
	})
}

func Increment(key string) {
	db := Open("my.db")

	currentCount := Get(db, "Video_game", key)

	currentCountInt, err := strconv.Atoi(currentCount)
	if err != nil {
		log.Fatal("Something's up with the database. Error converting to int.", err)
	}

	var newCount int
	newCount = currentCountInt + 1
	newCountStr := strconv.Itoa(newCount)

	Set(db, "Video_game", key, newCountStr)

	db.Close()
}

func Iterate(bucket string) ([]string, []opts.BarData) {
	values := make([]opts.BarData, 0)
	keys := make([]string, 0)

	db := Open("my.db")
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(bucket))

		c := b.Cursor()
		i := 0
		for k, v := c.First(); k != nil && i < 500; k, v = c.Next() {
			keys = append(keys, string(k))
			fmt.Println("key", i, ":", string(k))
			value, _ := strconv.Atoi(string(v))
			fmt.Println("val", i, ":", string(v))
			values = append(values, opts.BarData{Value: value})
			i++
		}

		return nil
	})
	db.Close()
	return keys, values
}

func GenGraph(category string) {
	keys, values := Iterate(category)

	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithXAxisOpts(opts.XAxis{
			AxisLabel: &opts.AxisLabel{
				Interval:     "0",
				Rotate:       30,
				ShowMinLabel: true,
				ShowMaxLabel: true,
				Color:        "#f5f5f5",
			},
		}),
		charts.WithYAxisOpts(opts.YAxis{
			AxisLabel: &opts.AxisLabel{
				Interval:     "0",
				ShowMinLabel: true,
				ShowMaxLabel: true,
				Color:        "#f5f5f5",
			},
			SplitLine: &opts.SplitLine{
				Show:      true,
				LineStyle: &opts.LineStyle{Color: "#f5f5f5"},
			},
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Width:           "1920px",
			Height:          "1080px",
			BackgroundColor: "transparent",
		}),
		charts.WithColorsOpts(opts.Colors{"#ef4444"}),
	)
	bar.SetXAxis(keys).AddSeries("Votes", values).SetSeriesOptions(charts.WithLineStyleOpts(opts.LineStyle{Color: "#262626"}))
	// bar.SetXAxis([]string{"1981 in video games", "1983 in video games", "Wed", "Thu", "Fri", "Sat", "Sun", "Sun", "Sun", "Sun"}).AddSeries("Votes", values)

	f, _ := os.Create("public/html/bar.html")
	bar.Render(f)
}
