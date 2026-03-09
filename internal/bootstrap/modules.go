package bootstrap

import (
	"fds-backend/internal/modulekit"
	"fds-backend/internal/modules/adminauth"
	designsubmissionmodule "fds-backend/internal/modules/designsubmission"
	"fds-backend/internal/modules/health"
	presetmodule "fds-backend/internal/modules/preset"
	"fds-backend/internal/modules/swagger"
)

func DefaultModules() []modulekit.Module {
	return []modulekit.Module{
		health.NewModule(),
		swagger.NewModule(),
		adminauth.NewModule(),
		designsubmissionmodule.NewModule(),
		presetmodule.NewModule(),
	}
}
