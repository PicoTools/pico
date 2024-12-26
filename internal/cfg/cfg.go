package cfg

import (
	"context"
	"errors"
	"os"

	"github.com/PicoTools/pico/internal/db"
	"github.com/PicoTools/pico/internal/listener"
	"github.com/PicoTools/pico/internal/management"
	"github.com/PicoTools/pico/internal/operator"
	"github.com/PicoTools/pico/internal/pki"

	"github.com/go-faster/sdk/zctx"
	"github.com/go-playground/validator/v10"
	"sigs.k8s.io/yaml"
)

// main config structure
type Config struct {
	Listener   listener.Config   `json:"listener" validate:"required"`
	Operator   operator.Config   `json:"operator" validate:"required"`
	Management management.Config `json:"management" validate:"required"`
	Db         db.Config         `json:"db" validate:"required"`
	Pki        pki.Config        `json:"pki"`
}

// Read get content of file by path and maps it on config structure
func Read(path string) (Config, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err = yaml.UnmarshalStrict(f, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

// Validate checks fileds in config structure
func (c Config) Validate(ctx context.Context) error {
	lg := zctx.From(ctx).Named("cfg-validator")

	v := validator.New(validator.WithRequiredStructEnabled())

	// register custom validator "one_of_nested"
	if err := v.RegisterValidation("one_of_nested", tagOneOfNested); err != nil {
		return err
	}
	// register custom validator "port"
	if err := v.RegisterValidation("port", tagPort); err != nil {
		return err
	}
	// register error translations
	if err := translations(v); err != nil {
		return err
	}
	// validate config
	if err := v.Struct(c); err != nil {
		var es validator.ValidationErrors
		if errors.As(err, &es) {
			for _, e := range es {
				lg.Error(e.Translate(trans))
			}
		}
		return errors.New("bunch of validation errors")
	}
	return nil
}
