package entities_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/valdirmendesdev/prayer-notebook/entities"
)

func TestNewPrayer(t *testing.T) {
	p, err := entities.NewPrayer("test", time.Now())
	assert.Nil(t, err)
	assert.NotNil(t, p.ID)
	assert.Equal(t, p.Talked, "test")
	assert.False(t, p.CreatedAt.IsZero())
	assert.True(t, p.UpdatedAt.IsZero())
}

func TestPrayer_Validate(t *testing.T) {
	testsCases := []struct {
		name   string
		talked string
		date   time.Time
	}{
		{
			name:   "All fields are empty",
			talked: "",
			date:   time.Time{},
		},
		{
			name:   "Talked field is empty",
			talked: "",
			date:   time.Now(),
		},
		{
			name:   "Date field is empty",
			talked: "test",
			date:   time.Time{},
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := entities.NewPrayer(tc.talked, tc.date)
			assert.NotNil(t, err)
		})
	}
}
