package cfg

import "github.com/spf13/viper"

func Load() error {
	viper.SetConfigFile("configs/application.yml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func MustGet[T any](key string) T {
	return viper.Get(key).(T)
}
