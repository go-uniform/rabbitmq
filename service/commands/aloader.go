package commands

var MustAsset func(string) []byte

func Load(mustAsset func(string) []byte) {
	// since all commands use init method to add themselves this function does nothing
	// calling this function will just the code optimizer from annoyingly removing the "unused" package import
	MustAsset = mustAsset
}