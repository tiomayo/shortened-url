package helper

import (
	"crypto/rand"
	"math"
	"math/big"
	"shortened-url/repository/routes"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const slugLength = 6

func GenerateSlug(repo routes.RepoInterface) string {
	for {
		shortKey := make([]byte, slugLength)
		for i := range shortKey {
			rand, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
			if err != nil {
				continue
			}
			shortKey[i] = charset[rand.Int64()]
		}

		slug := string(shortKey)
		if repo.Get(slug).Url == "" {
			return slug
		}
	}
}

func CalculateUniqueSlugs() int64 {
	return int64(math.Pow(float64(len(charset)), float64(slugLength)))
}
