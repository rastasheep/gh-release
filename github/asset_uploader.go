package github

type AssetUploader struct {
	Client  *Client
	Release *Release
}

func NewAssetUploader(c *Client, r *Release) *AssetUploader {
	return &AssetUploader{
		Client:  c,
		Release: r,
	}
}
