package node

type Node struct {
	Name            string
	IP              string
	Role            string
	Cores           int
	Memory          int
	MemoryAllocated int
	Disk            int
	DiskAllocated   int
	TaskCount       int
}
