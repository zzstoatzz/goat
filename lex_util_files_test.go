package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadPermissionedSpaceSchemas(t *testing.T) {
	tests := map[string]string{
		"space": `{
  "lexicon": 1,
  "id": "fm.plyr.privateMedia",
  "defs": {
    "main": {
      "type": "space",
      "key": "tid",
      "name": "Plyr private media",
      "collections": ["fm.plyr.track"]
    }
  }
}`,
		"permission set": `{
  "lexicon": 1,
  "id": "fm.plyr.privateMediaAccess",
  "defs": {
    "main": {
      "type": "permission-set",
      "title": "Private media access",
      "permissions": [
        {
          "type": "permission",
          "resource": "space",
          "spaceType": "fm.plyr.privateMedia",
          "authority": "self",
          "skey": "self",
          "collection": ["fm.plyr.track"],
          "action": ["read", "create", "update", "delete"],
          "manage": ["create", "update", "delete"]
        }
      ]
    }
  }
}`,
	}

	for name, schema := range tests {
		t.Run(name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "schema.json")
			if err := os.WriteFile(path, []byte(schema), 0o600); err != nil {
				t.Fatal(err)
			}

			if _, _, err := loadSchemaJSON(path); err != nil {
				t.Fatalf("load permissioned %s schema: %v", name, err)
			}
		})
	}
}
