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

//RESTInterfaceRouteTableUpdateRequest for update request for REST.
type RESTInterfaceRouteTableUpdateRequest struct {
    Data map[string]interface{} `json:"interface-route-table"`
}

//RESTCreateInterfaceRouteTable handle a Create REST service.
func (service *ContrailService) RESTCreateInterfaceRouteTable(c echo.Context) error {
    requestData := &models.InterfaceRouteTableCreateRequest{
        InterfaceRouteTable: models.MakeInterfaceRouteTable(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "interface_route_table",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateInterfaceRouteTable(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateInterfaceRouteTable handle a Create API
func (service *ContrailService) CreateInterfaceRouteTable(
    ctx context.Context, 
    request *models.InterfaceRouteTableCreateRequest) (*models.InterfaceRouteTableCreateResponse, error) {
    model := request.InterfaceRouteTable
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
            return db.CreateInterfaceRouteTable(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "interface_route_table",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.InterfaceRouteTableCreateResponse{
        InterfaceRouteTable: request.InterfaceRouteTable,
    }, nil
}

//RESTUpdateInterfaceRouteTable handles a REST Update request.
func (service *ContrailService) RESTUpdateInterfaceRouteTable(c echo.Context) error {
    id := c.Param("id")
    request := &models.InterfaceRouteTableUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "interface_route_table",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateInterfaceRouteTable(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateInterfaceRouteTable handles a Update request.
func (service *ContrailService) UpdateInterfaceRouteTable(ctx context.Context, request *models.InterfaceRouteTableUpdateRequest) (*models.InterfaceRouteTableUpdateResponse, error) {
    id = request.ID
    model = request.InterfaceRouteTable
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
            return db.UpdateInterfaceRouteTable(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "interface_route_table",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.InterfaceRouteTable.UpdateResponse{
        InterfaceRouteTable: model,
    }, nil
}

//RESTDeleteInterfaceRouteTable delete a resource using REST service.
func (service *ContrailService) RESTDeleteInterfaceRouteTable(c echo.Context) error {
    id := c.Param("id")
    request := &models.InterfaceRouteTableDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteInterfaceRouteTable(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteInterfaceRouteTable delete a resource.
func (service *ContrailService) DeleteInterfaceRouteTable(ctx context.Context, request *models.InterfaceRouteTableDeleteRequest) (*models.InterfaceRouteTableDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteInterfaceRouteTable(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.InterfaceRouteTableDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowInterfaceRouteTable a REST Show request.
func (service *ContrailService) RESTShowInterfaceRouteTable(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.InterfaceRouteTable
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListInterfaceRouteTable(tx, &common.ListSpec{
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
        "interface_route_table": result,
    })
}

//RESTListInterfaceRouteTable handles a List REST service Request.
func (service *ContrailService) RESTListInterfaceRouteTable(c echo.Context) (error) {
    var result []*models.InterfaceRouteTable
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListInterfaceRouteTable(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "interface-route-tables": result,
    })
}