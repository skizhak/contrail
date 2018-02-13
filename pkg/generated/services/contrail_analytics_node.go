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

//RESTContrailAnalyticsNodeUpdateRequest for update request for REST.
type RESTContrailAnalyticsNodeUpdateRequest struct {
    Data map[string]interface{} `json:"contrail-analytics-node"`
}

//RESTCreateContrailAnalyticsNode handle a Create REST service.
func (service *ContrailService) RESTCreateContrailAnalyticsNode(c echo.Context) error {
    requestData := &models.ContrailAnalyticsNodeCreateRequest{
        ContrailAnalyticsNode: models.MakeContrailAnalyticsNode(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "contrail_analytics_node",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateContrailAnalyticsNode(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateContrailAnalyticsNode handle a Create API
func (service *ContrailService) CreateContrailAnalyticsNode(
    ctx context.Context, 
    request *models.ContrailAnalyticsNodeCreateRequest) (*models.ContrailAnalyticsNodeCreateResponse, error) {
    model := request.ContrailAnalyticsNode
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
            return db.CreateContrailAnalyticsNode(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "contrail_analytics_node",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.ContrailAnalyticsNodeCreateResponse{
        ContrailAnalyticsNode: request.ContrailAnalyticsNode,
    }, nil
}

//RESTUpdateContrailAnalyticsNode handles a REST Update request.
func (service *ContrailService) RESTUpdateContrailAnalyticsNode(c echo.Context) error {
    id := c.Param("id")
    request := &models.ContrailAnalyticsNodeUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "contrail_analytics_node",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateContrailAnalyticsNode(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateContrailAnalyticsNode handles a Update request.
func (service *ContrailService) UpdateContrailAnalyticsNode(ctx context.Context, request *models.ContrailAnalyticsNodeUpdateRequest) (*models.ContrailAnalyticsNodeUpdateResponse, error) {
    id = request.ID
    model = request.ContrailAnalyticsNode
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
            return db.UpdateContrailAnalyticsNode(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "contrail_analytics_node",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.ContrailAnalyticsNode.UpdateResponse{
        ContrailAnalyticsNode: model,
    }, nil
}

//RESTDeleteContrailAnalyticsNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteContrailAnalyticsNode(c echo.Context) error {
    id := c.Param("id")
    request := &models.ContrailAnalyticsNodeDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteContrailAnalyticsNode(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteContrailAnalyticsNode delete a resource.
func (service *ContrailService) DeleteContrailAnalyticsNode(ctx context.Context, request *models.ContrailAnalyticsNodeDeleteRequest) (*models.ContrailAnalyticsNodeDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteContrailAnalyticsNode(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.ContrailAnalyticsNodeDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowContrailAnalyticsNode a REST Show request.
func (service *ContrailService) RESTShowContrailAnalyticsNode(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.ContrailAnalyticsNode
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListContrailAnalyticsNode(tx, &common.ListSpec{
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
        "contrail_analytics_node": result,
    })
}

//RESTListContrailAnalyticsNode handles a List REST service Request.
func (service *ContrailService) RESTListContrailAnalyticsNode(c echo.Context) (error) {
    var result []*models.ContrailAnalyticsNode
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListContrailAnalyticsNode(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "contrail-analytics-nodes": result,
    })
}