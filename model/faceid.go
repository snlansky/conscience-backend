package model

type FaceID struct {
	ID         string                 `json:"id" `          // ID
	SourceType string                 `json:"source_type" ` // 资源类型
	SourceHash string                 `json:"source_hash" ` // 资源hash
	Algorithm  string                 `json:"algorithm" `   // hash 算法
	Labels     []string               `json:"labels" `      // 标签
	Metadata   map[string]interface{} `json:"metadata" `    // 元数据
	Timestamp  int64                  `json:"timestamp"`    // 时间戳(s)
}
