package manager

import (
	"errors"
	"password-manager/storage"
	"password-manager/utils"
)

type Manager struct {
	store storage.Storage
}

func NewManager(path string) *Manager {
	return &Manager{
		store: storage.NewFileStorage(path),
	}
}

// Add password to file
func (m *Manager) Add(name, plainPassword string) error {
	entries, err := m.store.Load()
	if err != nil {
		return err
	}
	encrypted, err := utils.Encrypt(plainPassword)
	if err != nil {
		return err
	}
	entries = append(entries, storage.Entry{Name: name, Password: encrypted})
	return m.store.Save(entries)
}

// Generate password to file
func (m *Manager) Generate(name string, length int) (string, error) {
	pwd := utils.Generate(length)
	err := m.Add(name, pwd)
	return pwd, err
}

// Delete password from file
func (m *Manager) Delete(name string) error {
	entries, err := m.store.Load()
	if err != nil {
		return err
	}
	found := false
	var filtered []storage.Entry
	for _, e := range entries {
		if e.Name != name {
			filtered = append(filtered, e)
		} else {
			found = true
		}
	}
	if !found {
		return errors.New("not found")
	}
	return m.store.Save(filtered)
}

// List of passwords
func (m *Manager) List() []storage.Entry {
	entries, err := m.store.Load()
	if err != nil {
		return []storage.Entry{}
	}
	for i, e := range entries {
		decrypted, err := utils.Decrypt(e.Password)
		if err == nil {
			entries[i].Password = decrypted
		}
	}
	return entries
}
