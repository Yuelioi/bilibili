package login

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// 提取Keys里的hash值
func extract(url string) string {

	parts := strings.Split(url, "/")
	filename := parts[len(parts)-1]
	hash := strings.Split(filename, ".")[0]
	return hash
}

var (
	mixinKeyEncTab = []int{
		46, 47, 18, 2, 53, 8, 23, 32, 15, 50, 10, 31, 58, 3, 45, 35, 27, 43, 5, 49,
		33, 9, 42, 19, 29, 28, 14, 39, 12, 38, 41, 13, 37, 48, 7, 16, 24, 55, 40,
		61, 26, 17, 0, 1, 60, 51, 30, 4, 22, 25, 54, 21, 56, 59, 6, 63, 57, 62, 11,
		36, 20, 34, 44, 52,
	}
	cache          sync.Map
	lastUpdateTime time.Time
)

func encWbi(params map[string]string, imgKey, subKey string) map[string]string {
	mixinKey := getMixinKey(imgKey + subKey)
	currTime := strconv.FormatInt(time.Now().Unix(), 10)
	params["wts"] = currTime

	// Sort keys
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Remove unwanted characters
	for k, v := range params {
		v = sanitizeString(v)
		params[k] = v
	}

	// Build URL parameters
	query := url.Values{}
	for _, k := range keys {
		query.Set(k, params[k])
	}
	queryStr := query.Encode()

	// Calculate w_rid
	hash := md5.Sum([]byte(queryStr + mixinKey))
	params["w_rid"] = hex.EncodeToString(hash[:])
	return params
}

func getMixinKey(orig string) string {
	var str strings.Builder
	for _, v := range mixinKeyEncTab {
		if v < len(orig) {
			str.WriteByte(orig[v])
		}
	}
	return str.String()[:32]
}

func sanitizeString(s string) string {
	unwantedChars := []string{"!", "'", "(", ")", "*"}
	for _, char := range unwantedChars {
		s = strings.ReplaceAll(s, char, "")
	}
	return s
}

func updateCache(a *Login) {
	if time.Since(lastUpdateTime).Minutes() < 10 {
		return
	}
	keys, err := a.UserKeys()

	if err != nil {
		fmt.Printf("keys 获取失败, err:%s", err)
		return
	}
	imgKey := keys.ImgURL
	subKey := keys.SubURL
	cache.Store("imgKey", imgKey)
	cache.Store("subKey", subKey)
	lastUpdateTime = time.Now()
}

func getWbiKeysCached(a *Login) (string, string) {
	updateCache(a)
	imgKeyI, _ := cache.Load("imgKey")
	subKeyI, _ := cache.Load("subKey")
	return imgKeyI.(string), subKeyI.(string)
}
