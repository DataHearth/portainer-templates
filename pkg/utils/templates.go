package utils

import (
	"encoding/json"
	"sync"

	"github.com/datahearth/portainer-templates/pkg/db/tables"
	"github.com/datahearth/portainer-templates/pkg/db/templates"
)

// todo: add waitgroups for each func
// todo: delete data after being used in loops
func FormatBody(containers []tables.ContainerTable, stacks []tables.StackTable, composes []tables.ComposeTable) ([]byte, error) {
	t := new(templates.Templates)
	t.Version = "2"

	for _, c := range composes {
		t.Templates = append(t.Templates, SQLComposeToJSON(c))
	}
	for _, c := range containers {
		t.Templates = append(t.Templates, SQLContainerToJSON(c))
	}
	for _, s := range stacks {
		t.Templates = append(t.Templates, SQLStackToJSON(s))
	}

	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func JSONStackToSQL(stack templates.Stack) tables.StackTable {
	wg := new(sync.WaitGroup)
	cats := make([]tables.StackCategory, 0, len(stack.Categories))
	envs := make([]tables.StackEnv, 0, len(stack.Env))
	repo := tables.StackRepository{
		ID:        stack.Repository.ID,
		URL:       stack.Repository.URL,
		Stackfile: stack.Repository.Stackfile,
	}

	wg.Add(2)
	go func() {
		for _, cat := range stack.Categories {
			cats = append(cats, tables.StackCategory{
				Name: cat,
			})
		}
		wg.Done()
	}()
	go func() {
		for _, e := range stack.Env {
			selects := make([]tables.StackSelect, 0, len(e.Select))

			for _, s := range e.Select {
				selects = append(selects, tables.StackSelect{
					ID:      s.ID,
					Text:    s.Text,
					Value:   s.Value,
					Default: s.Default,
				})
			}

			envs = append(envs, tables.StackEnv{
				ID:          e.ID,
				Name:        e.Name,
				Label:       e.Label,
				Description: e.Description,
				Default:     e.Default,
				Preset:      e.Preset,
				Selects:     selects,
			})
		}
		wg.Done()
	}()

	wg.Wait()

	return tables.StackTable{
		ID:                stack.ID,
		Type:              stack.Type,
		Title:             stack.Title,
		Description:       stack.Description,
		Note:              stack.Note,
		Categories:        cats,
		Platform:          stack.Platform,
		Logo:              stack.Logo,
		Repository:        repo,
		Envs:              envs,
		AdministratorOnly: stack.AdministratorOnly,
		Name:              stack.Name,
	}
}

func JSONContainerToSQL(container templates.Container) tables.ContainerTable {
	wg := new(sync.WaitGroup)
	cats := make([]tables.ContainerCategory, 0, len(container.Categories))
	ports := make([]tables.ContainerPort, 0, len(container.Ports))
	volumes := make([]tables.ContainerVolume, 0, len(container.Volumes))
	envs := make([]tables.ContainerEnv, 0, len(container.Env))
	labels := make([]tables.ContainerLabel, 0, len(container.Labels))

	wg.Add(5)
	go func() {
		for _, c := range container.Categories {
			cats = append(cats, tables.ContainerCategory{
				Name: c,
			})
		}
		wg.Done()
	}()
	go func() {
		for _, p := range container.Ports {
			ports = append(ports, tables.ContainerPort{
				Port: p,
			})
		}
		wg.Done()
	}()
	go func() {
		for _, v := range container.Volumes {
			volumes = append(volumes, tables.ContainerVolume{
				Bind:      v.Bind,
				Container: v.Container,
				ReadOnly:  v.ReadOnly,
			})
		}
		wg.Done()
	}()
	go func() {
		for _, e := range container.Env {
			selects := make([]tables.ContainerSelect, 0, len(e.Select))

			for _, s := range e.Select {
				selects = append(selects, tables.ContainerSelect{
					Text:    s.Text,
					Value:   s.Value,
					Default: s.Default,
				})
			}

			envs = append(envs, tables.ContainerEnv{
				Name:        e.Name,
				Label:       e.Label,
				Default:     e.Default,
				Description: e.Description,
				Preset:      e.Preset,
				Selects:     selects,
			})
		}
		wg.Done()
	}()
	go func() {
		for _, l := range container.Labels {
			labels = append(labels, tables.ContainerLabel{
				Name:  l.Name,
				Value: l.Value,
			})
		}
		wg.Done()
	}()

	wg.Wait()

	return tables.ContainerTable{
		Type:              container.Type,
		Title:             container.Title,
		Description:       container.Description,
		Categories:        cats,
		Platform:          container.Platform,
		Logo:              container.Logo,
		Image:             container.Image,
		Ports:             ports,
		Volumes:           volumes,
		AdministratorOnly: container.AdministratorOnly,
		Name:              container.Name,
		Registry:          container.Registry,
		Command:           container.Command,
		Envs:              envs,
		Network:           container.Network,
		Labels:            labels,
		Privileged:        container.Privileged,
		Interactive:       container.Interactive,
		RestartPolicy:     container.RestartPolicy,
		Hostname:          container.Hostname,
		Note:              container.Note,
	}
}

func JSONComposeToSQL(compose templates.Compose) tables.ComposeTable {
	wg := new(sync.WaitGroup)
	cats := make([]tables.ComposeCategory, 0, len(compose.Categories))
	envs := make([]tables.ComposeEnv, 0, len(compose.Env))
	repo := tables.ComposeRepository{
		ID:        compose.Repository.ID,
		URL:       compose.Repository.URL,
		Stackfile: compose.Repository.Stackfile,
	}

	wg.Add(2)
	go func() {
		for _, cat := range compose.Categories {
			cats = append(cats, tables.ComposeCategory{
				Name: cat,
			})
		}
		wg.Done()
	}()
	go func() {
		for _, e := range compose.Env {
			selects := make([]tables.ComposeSelect, 0, len(e.Select))

			for _, s := range e.Select {
				selects = append(selects, tables.ComposeSelect{
					ID:      s.ID,
					Text:    s.Text,
					Value:   s.Value,
					Default: s.Default,
				})
			}

			envs = append(envs, tables.ComposeEnv{
				ID:          e.ID,
				Name:        e.Name,
				Label:       e.Label,
				Description: e.Description,
				Default:     e.Default,
				Preset:      e.Preset,
				Selects:     selects,
			})
		}
		wg.Done()
	}()

	wg.Wait()

	return tables.ComposeTable{
		ID:                compose.ID,
		Type:              compose.Type,
		Title:             compose.Title,
		Description:       compose.Description,
		Note:              compose.Note,
		Categories:        cats,
		Platform:          compose.Platform,
		Logo:              compose.Logo,
		Repository:        repo,
		Envs:              envs,
		AdministratorOnly: compose.AdministratorOnly,
		Name:              compose.Name,
	}
}

func SQLStackToJSON(stack tables.StackTable) templates.Stack {
	wg := new(sync.WaitGroup)
	cats := make([]string, 0, len(stack.Categories))
	envs := make([]templates.Env, 0, len(stack.Envs))
	repo := templates.Repository{
		ID:        stack.Repository.ID,
		URL:       stack.Repository.URL,
		Stackfile: stack.Repository.Stackfile,
	}

	wg.Add(2)
	go func() {
		for _, c := range stack.Categories {
			cats = append(cats, c.Name)
		}
		wg.Done()
	}()
	go func() {
		for _, e := range stack.Envs {
			selects := make([]templates.Select, 0, len(e.Selects))

			for _, s := range e.Selects {
				selects = append(selects, templates.Select{
					ID:      s.ID,
					Text:    s.Text,
					Value:   s.Value,
					Default: s.Default,
				})
			}

			envs = append(envs, templates.Env{
				ID:          e.ID,
				Name:        e.Name,
				Label:       e.Label,
				Description: e.Description,
				Default:     e.Default,
				Preset:      e.Preset,
				Select:      selects,
			})
		}
		wg.Done()
	}()

	wg.Wait()

	return templates.Stack{
		ID:                stack.ID,
		Type:              stack.Type,
		Title:             stack.Title,
		Description:       stack.Description,
		Note:              stack.Note,
		Categories:        cats,
		Platform:          stack.Platform,
		Logo:              stack.Logo,
		Repository:        repo,
		Env:               envs,
		AdministratorOnly: stack.AdministratorOnly,
		Name:              stack.Name,
	}
}

func SQLContainerToJSON(container tables.ContainerTable) templates.Container {
	wg := new(sync.WaitGroup)
	cats := make([]string, 0, len(container.Categories))
	envs := make([]templates.Env, 0, len(container.Envs))
	ports := make([]string, 0, len(container.Ports))
	volumes := make([]templates.Volumes, 0, len(container.Volumes))
	labels := make([]templates.Label, 0, len(container.Labels))

	wg.Add(5)
	go func() {
		for _, cat := range container.Categories {
			cats = append(cats, cat.Name)
		}
		wg.Done()
	}()
	go func() {
		for _, e := range container.Envs {
			selects := make([]templates.Select, 0, len(e.Selects))

			for _, s := range e.Selects {
				selects = append(selects, templates.Select{
					ID:      s.ID,
					Text:    s.Text,
					Value:   s.Value,
					Default: s.Default,
				})
			}

			envs = append(envs, templates.Env{
				ID:          e.ID,
				Name:        e.Name,
				Label:       e.Label,
				Description: e.Description,
				Default:     e.Default,
				Preset:      e.Preset,
				Select:      selects,
			})
		}
		wg.Done()
	}()
	go func() {
		for _, p := range container.Ports {
			ports = append(ports, p.Port)
		}
		wg.Done()
	}()
	go func() {
		for _, v := range container.Volumes {
			volumes = append(volumes, templates.Volumes{
				ID:        v.ID,
				Container: v.Container,
				Bind:      v.Bind,
				ReadOnly:  v.ReadOnly,
			})
		}
		wg.Done()
	}()
	go func() {
		for _, l := range container.Labels {
			labels = append(labels, templates.Label{
				ID:    l.ID,
				Name:  l.Name,
				Value: l.Value,
			})
		}
		wg.Done()
	}()

	wg.Wait()

	return templates.Container{
		ID:                container.ID,
		Type:              container.Type,
		Title:             container.Title,
		Description:       container.Description,
		Note:              container.Note,
		Categories:        cats,
		Platform:          container.Platform,
		Logo:              container.Logo,
		Env:               envs,
		AdministratorOnly: container.AdministratorOnly,
		Name:              container.Name,
		Image:             container.Image,
		Ports:             ports,
		Volumes:           volumes,
		Registry:          container.Registry,
		Command:           container.Command,
		Network:           container.Network,
		Labels:            labels,
		Privileged:        container.Privileged,
		Interactive:       container.Interactive,
		RestartPolicy:     container.RestartPolicy,
		Hostname:          container.Hostname,
	}
}

func SQLComposeToJSON(compose tables.ComposeTable) templates.Compose {
	wg := new(sync.WaitGroup)
	cats := make([]string, 0, len(compose.Categories))
	envs := make([]templates.Env, 0, len(compose.Envs))
	repo := templates.Repository{
		ID:        compose.Repository.ID,
		URL:       compose.Repository.URL,
		Stackfile: compose.Repository.Stackfile,
	}

	wg.Add(2)
	go func() {
		for _, cat := range compose.Categories {
			cats = append(cats, cat.Name)
		}
		wg.Done()
	}()
	go func() {
		for _, e := range compose.Envs {
			selects := make([]templates.Select, 0, len(e.Selects))

			for _, s := range e.Selects {
				selects = append(selects, templates.Select{
					ID:      s.ID,
					Text:    s.Text,
					Value:   s.Value,
					Default: s.Default,
				})
			}

			envs = append(envs, templates.Env{
				ID:          e.ID,
				Name:        e.Name,
				Label:       e.Label,
				Description: e.Description,
				Default:     e.Default,
				Preset:      e.Preset,
				Select:      selects,
			})
		}
		wg.Done()
	}()

	wg.Wait()

	return templates.Compose{
		ID:                compose.ID,
		Type:              compose.Type,
		Title:             compose.Title,
		Description:       compose.Description,
		Note:              compose.Note,
		Categories:        cats,
		Platform:          compose.Platform,
		Logo:              compose.Logo,
		Repository:        repo,
		Env:               envs,
		AdministratorOnly: compose.AdministratorOnly,
		Name:              compose.Name,
	}
}
