/*
* Our database server is too slow.
* We are going to create a custom memory cache. We’ll use Go’s map collection type, which will act as the cache.
* There is a global limit on the number of items that can be in the cache. We’ll use one map to help keep track of the number of items in the cache.
* We have two types of data we need to cache: books and CDs.
* Both use the ID, so we need a way to separate the two types of items in the shared cache. (prefixes)
* We need away to set and get items from the cache.
 */

package main

import "fmt"

const GlobalLimit = 100

// create a max cache size that is 10 times the global limit
const MaxCacheSize int = 10 * GlobalLimit

// create cache prefixes
const (
	CacheKeyBook = "book_"
	CacheKeyCD   = "cd_"
)

// map value that has a string value for a key, and a string value foor its values as my cache
var cache map[string]string

// get items from the cache
func cacheGet(key string) string {
	return cache[key]
}

// set itesms in the cache
func cacheSet(key, val string) {
	if len(cache)+1 >= MaxCacheSize {
		return
	}
	cache[key] = val
}

// get a book from a cache
func GetBook(isbn string) string {
	return cacheGet(CacheKeyBook + isbn)
}

// add a book to the cache
func SetBook(isbn string, name string) {
	cacheSet(CacheKeyBook+isbn, name)
}

// get a CD data from the cache
func GetCD(sku string) string {
	return cacheGet(CacheKeyCD + sku)
}

// add CDs to the cache
func SetCD(sku string, title string) {
	cacheSet(CacheKeyCD+sku, title)
}

func main() {
	// initialize cache as a map
	cache = make(map[string]string)

	// add a book and a CD to the cache
	SetBook("1234-5678", "Get Ready To Go")
	SetCD("1234-5678", "Get Ready To Go Audio Book")

	// get and print book and CD
	fmt.Println("Book :", GetBook("1234-5678"))
	fmt.Println("CD :", GetCD("1234-5678"))
}
