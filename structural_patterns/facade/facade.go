package facade

import "fmt"

// Subsystem classes
type ComputeManager struct{}

func (c *ComputeManager) StartInstance(instanceID string) {
	fmt.Printf("Starting instance %s.\n", instanceID)
}

func (c *ComputeManager) StopInstance(instanceID string) {
	fmt.Printf("Stopping instance %s.\n", instanceID)
}

type StorageManager struct{}

func (s *StorageManager) CreateVolume(sizeGB int, volumeType string) {
	fmt.Printf("Creating a %d GB %s volume.\n", sizeGB, volumeType)
}

func (s *StorageManager) DeleteVolume(volumeID string) {
	fmt.Printf("Deleting volume %s.\n", volumeID)
}

type NetworkManager struct{}

func (n *NetworkManager) CreateVPC(cidr string) {
	fmt.Printf("Creating VPC with CIDR %s.\n", cidr)
}

func (n *NetworkManager) DeleteVPC(vpcID string) {
	fmt.Printf("Deleting VPC %s.\n", vpcID)
}

// Facade
type CloudFacade struct {
	compute *ComputeManager
	storage *StorageManager
	network *NetworkManager
}

func NewCloudFacade() *CloudFacade {
	return &CloudFacade{
		compute: &ComputeManager{},
		storage: &StorageManager{},
		network: &NetworkManager{},
	}
}

func (f *CloudFacade) DeployApplication() {
	fmt.Println("Deploying application...")
	f.network.CreateVPC("192.168.0.0/16")
	f.compute.StartInstance("i-123456")
	f.storage.CreateVolume(100, "SSD")
}

func (f *CloudFacade) TeardownApplication() {
	fmt.Println("Tearing down application...")
	f.compute.StopInstance("i-123456")
	f.storage.DeleteVolume("vol-123456")
	f.network.DeleteVPC("vpc-123456")
}
