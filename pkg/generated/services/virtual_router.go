package services 

import (
    "context"
    "net/http"
    "database/sql"
    "github.com/Juniper/contrail/pkg/generated/models"
    "github.com/Juniper/contrail/pkg/generated/db"
    "github.com/satori/go.uuid"
    "github.com/labstack/echo"
    "github.com/Juniper/contrail/pkg/common"

	log "github.com/sirupsen/logrus"
)

//RESTVirtualRouterUpdateRequest for update request for REST.
type RESTVirtualRouterUpdateRequest struct {
    Data map[string]interface{} `json:"virtual-router"`
}

//RESTCreateVirtualRouter handle a Create REST service.
func (service *ContrailService) RESTCreateVirtualRouter(c echo.Context) error {
    requestData := &models.VirtualRouterCreateRequest{
        VirtualRouter: models.MakeVirtualRouter(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "virtual_router",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateVirtualRouter(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateVirtualRouter handle a Create API
func (service *ContrailService) CreateVirtualRouter(
    ctx context.Context, 
    request *models.VirtualRouterCreateRequest) (*models.VirtualRouterCreateResponse, error) {
    model := request.VirtualRouter
    if model.UUID == "" {
        model.UUID = uuid.NewV4().String()
    }

    if model.FQName == nil {
       return nil, common.ErrorBadRequest("Missing fq_name")
    }

    auth := common.GetAuthCTX(ctx)
    if auth == nil {
        return nil, common.ErrorUnauthenticated
    }
    model.Perms2.Owner = auth.ProjectID()
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.CreateVirtualRouter(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "virtual_router",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.VirtualRouterCreateResponse{
        VirtualRouter: request.VirtualRouter,
    }, nil
}

//RESTUpdateVirtualRouter handles a REST Update request.
func (service *ContrailService) RESTUpdateVirtualRouter(c echo.Context) error {
    id := c.Param("id")
    request := &models.VirtualRouterUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "virtual_router",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateVirtualRouter(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateVirtualRouter handles a Update request.
func (service *ContrailService) UpdateVirtualRouter(ctx context.Context, request *models.VirtualRouterUpdateRequest) (*models.VirtualRouterUpdateResponse, error) {
    id = request.ID
    model = request.VirtualRouter
    if model == nil {
        return nil, common.ErrorBadRequest("Update body is empty")
    }
    auth := common.GetAuthCTX(ctx)
    ok := common.SetValueByPath(model, "Perms2.Owner", ".", auth.ProjectID())
    if !ok {
        return nil, common.ErrorBadRequest("Invalid JSON format")
    }
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.UpdateVirtualRouter(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "virtual_router",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.VirtualRouter.UpdateResponse{
        VirtualRouter: model,
    }, nil
}

//RESTDeleteVirtualRouter delete a resource using REST service.
func (service *ContrailService) RESTDeleteVirtualRouter(c echo.Context) error {
    id := c.Param("id")
    request := &models.VirtualRouterDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteVirtualRouter(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteVirtualRouter delete a resource.
func (service *ContrailService) DeleteVirtualRouter(ctx context.Context, request *models.VirtualRouterDeleteRequest) (*models.VirtualRouterDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteVirtualRouter(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.VirtualRouterDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowVirtualRouter a REST Show request.
func (service *ContrailService) RESTShowVirtualRouter(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.VirtualRouter
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListVirtualRouter(tx, &common.ListSpec{
                Limit: 1,
                Auth: auth,
                Filter: common.Filter{
                    "uuid": []string{id},
                },
            })
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "virtual_router": result,
    })
}

//RESTListVirtualRouter handles a List REST service Request.
func (service *ContrailService) RESTListVirtualRouter(c echo.Context) (error) {
    var result []*models.VirtualRouter
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListVirtualRouter(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "virtual-routers": result,
    })
}