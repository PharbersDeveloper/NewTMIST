package NtmResource

import (
	"errors"
	"Ntm/NtmDataStorage"
	"Ntm/NtmModel"
	"reflect"
	"net/http"

	"github.com/alfredyang1986/BmServiceDef/BmDataStorage"
	"github.com/manyminds/api2go"
)

type NtmLevelConfigResource struct {
	NtmLevelConfigStorage          		*NtmDataStorage.NtmLevelConfigStorage
	NtmRegionalDivisionResultStorage	*NtmDataStorage.NtmRegionalDivisionResultStorage
	NtmTargetAssignsResultStorage		*NtmDataStorage.NtmTargetAssignsResultStorage
	NtmResourceAssignsResultStorage		*NtmDataStorage.NtmResourceAssignsResultStorage
	NtmManageTimeResultStorage			*NtmDataStorage.NtmManageTimeResultStorage
	NtmManageTeamResultStorage			*NtmDataStorage.NtmManageTeamResultStorage
}

func (c NtmLevelConfigResource) NewLevelConfigResource(args []BmDataStorage.BmStorage) *NtmLevelConfigResource {
	var lcs *NtmDataStorage.NtmLevelConfigStorage
	var rdr	*NtmDataStorage.NtmRegionalDivisionResultStorage
	var tar	*NtmDataStorage.NtmTargetAssignsResultStorage
	var rar *NtmDataStorage.NtmResourceAssignsResultStorage
	var mtr *NtmDataStorage.NtmManageTimeResultStorage
	var mtrs *NtmDataStorage.NtmManageTeamResultStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmLevelConfigStorage" {
			lcs = arg.(*NtmDataStorage.NtmLevelConfigStorage)
		} else if tp.Name() == "NtmRegionalDivisionResultStorage" {
			rdr = arg.(*NtmDataStorage.NtmRegionalDivisionResultStorage)
		} else if tp.Name() == "NtmTargetAssignsResultStorage" {
			tar = arg.(*NtmDataStorage.NtmTargetAssignsResultStorage)
		} else if tp.Name() == "NtmResourceAssignsResultStorage" {
			rar = arg.(*NtmDataStorage.NtmResourceAssignsResultStorage)
		} else if tp.Name() == "NtmManageTimeResultStorage" {
			mtr = arg.(*NtmDataStorage.NtmManageTimeResultStorage)
		} else if tp.Name() == "NtmManageTeamResultStorage" {
			mtrs = arg.(*NtmDataStorage.NtmManageTeamResultStorage)
		}
	}
	return &NtmLevelConfigResource{
		NtmLevelConfigStorage:	lcs,
		NtmRegionalDivisionResultStorage: rdr,
		NtmTargetAssignsResultStorage: tar,
		NtmResourceAssignsResultStorage: rar,
		NtmManageTimeResultStorage: mtr,
		NtmManageTeamResultStorage: mtrs,
	}
}

// FindAll images
func (c NtmLevelConfigResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []NtmModel.LevelConfig
	regionalDivisionResultsID, rdrOk := r.QueryParams["regionalDivisionResultsID"]
	targetAssignsResultsID, tarOk := r.QueryParams["targetAssignsResultsID"]
	resourceAssignsResultsID, rarOk := r.QueryParams["resourceAssignsResultsID"]
	manageTimeResultsID, mtrOk := r.QueryParams["manageTimeResultsID"]
	manageTeamResultsID, mtrsOk := r.QueryParams["manageTeamResultsID"]

	if rdrOk {
		modelRootID := regionalDivisionResultsID[0]
		modelRoot, err := c.NtmRegionalDivisionResultStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= c.NtmLevelConfigStorage.GetOne(modelRoot.LevelConfigID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	if tarOk {
		modelRootID := targetAssignsResultsID[0]
		modelRoot, err := c.NtmTargetAssignsResultStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= c.NtmLevelConfigStorage.GetOne(modelRoot.LevelConfigID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	if rarOk {
		modelRootID := resourceAssignsResultsID[0]
		modelRoot, err := c.NtmResourceAssignsResultStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= c.NtmLevelConfigStorage.GetOne(modelRoot.LevelConfigID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	if mtrOk {
		modelRootID := manageTimeResultsID[0]
		modelRoot, err := c.NtmManageTimeResultStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= c.NtmLevelConfigStorage.GetOne(modelRoot.LevelConfigID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	if mtrsOk {
		modelRootID := manageTeamResultsID[0]
		modelRoot, err := c.NtmManageTeamResultStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= c.NtmLevelConfigStorage.GetOne(modelRoot.LevelConfigID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	result = c.NtmLevelConfigStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c NtmLevelConfigResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.NtmLevelConfigStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c NtmLevelConfigResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.LevelConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.NtmLevelConfigStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c NtmLevelConfigResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.NtmLevelConfigStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c NtmLevelConfigResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.LevelConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.NtmLevelConfigStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
