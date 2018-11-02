package vault

import (
	"github.com/hashicorp/vault/api"
	"github.com/jenkins-x/jx/pkg/kube"
)

// Vaulter is an interface for creating new vault clients
type Vaulter interface {
	NewVaultClient() api.Client
}

// VaultSelector is an interface for selecting a vault from the installed ones on the platform
// It should pick the most logical one, or give the user a way of picking a vault if there are multiple installed
type VaultSelector interface {
	GetVault(namespaces string) (*kube.Vault, error)
}
