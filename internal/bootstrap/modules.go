package bootstrap

import (
	"fds-backend/internal/modulekit"
	"fds-backend/internal/modules/adminauth"
	designsubmissionmodule "fds-backend/internal/modules/designsubmission"
	fileassetmodule "fds-backend/internal/modules/fileasset"
	"fds-backend/internal/modules/health"
	presetmodule "fds-backend/internal/modules/preset"
	"fds-backend/internal/modules/swagger"
)

func DefaultModules() []modulekit.Module {
	return []modulekit.Module{
		health.NewModule(),
		swagger.NewModule(),
		adminauth.NewModule(),
		fileassetmodule.NewModule(),
		designsubmissionmodule.NewModule(),
		presetmodule.NewModule(),
	}
}
