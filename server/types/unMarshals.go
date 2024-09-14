package types

import (
	"encoding/json"
	// "log"
)


func (n *MALDataNode) UnmarshalJSON(data []byte) error {
    // Create a temporary struct to capture known fields
    type Alias MALDataNode
    aux := &struct {
        *Alias
        AdditionalFields map[string]json.RawMessage `json:"-"` // To capture dynamic fields
    }{
        Alias: (*Alias)(n),
        AdditionalFields: make(map[string]json.RawMessage),
    }

    // Unmarshal known fields
    if err := json.Unmarshal(data, &aux); err != nil {
        return err
    }

    // Re-marshal and unmarshal again to capture unknown fields
    var rawMap map[string]json.RawMessage
    if err := json.Unmarshal(data, &rawMap); err != nil {
        return err
    }

    // Populate CustomFields with dynamic fields (excluding known ones)
    n.CustomFields = make(map[string]interface{})
    for key, value := range rawMap {
        if key != "id" && key != "main_picture" && key != "title" {
            var fieldValue interface{}
            if err := json.Unmarshal(value, &fieldValue); err == nil {
                n.CustomFields[key] = fieldValue
            } else {
                n.CustomFields[key] = string(value) // Fallback to string
            }
        }
    }

    return nil
}
