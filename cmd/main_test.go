package main

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnmarshalJSON(t *testing.T) {
	var d JSONDashboard
	f, err := ioutil.ReadFile("files/dashboard.json")
	require.NoError(t, err)

	err = json.Unmarshal(f, &d)
	require.NoError(t, err)
}
