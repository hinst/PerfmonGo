package perfmongo

import "sync"

var CachedAssets map[string][]byte = make(map[string][]byte)
var CachedAssetsLocker sync.Mutex

func GetCachedAsset(name string) []byte {
	CachedAssetsLocker.Lock()
	defer CachedAssetsLocker.Unlock()
	var value, success = CachedAssets[name]
	if success {
		return value
	} else {
		value, _ = GetAsset(name)
		CachedAssets[name] = value
	}
	return value
}
