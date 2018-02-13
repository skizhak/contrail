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

//RESTKubernetesClusterUpdateRequest for update request for REST.
type RESTKubernetesClusterUpdateRequest struct {
    Data map[string]interface{} `json:"kubernetes-cluster"`
}

//RESTCreateKubernetesCluster handle a Create REST service.
func (service *ContrailService) RESTCreateKubernetesCluster(c echo.Context) error {
    requestData := &models.KubernetesClusterCreateRequest{
        KubernetesCluster: models.MakeKubernetesCluster(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "kubernetes_cluster",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateKubernetesCluster(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateKubernetesCluster handle a Create API
func (service *ContrailService) CreateKubernetesCluster(
    ctx context.Context, 
    request *models.KubernetesClusterCreateRequest) (*models.KubernetesClusterCreateResponse, error) {
    model := request.KubernetesCluster
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
            return db.CreateKubernetesCluster(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "kubernetes_cluster",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.KubernetesClusterCreateResponse{
        KubernetesCluster: request.KubernetesCluster,
    }, nil
}

//RESTUpdateKubernetesCluster handles a REST Update request.
func (service *ContrailService) RESTUpdateKubernetesCluster(c echo.Context) error {
    id := c.Param("id")
    request := &models.KubernetesClusterUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "kubernetes_cluster",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateKubernetesCluster(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateKubernetesCluster handles a Update request.
func (service *ContrailService) UpdateKubernetesCluster(ctx context.Context, request *models.KubernetesClusterUpdateRequest) (*models.KubernetesClusterUpdateResponse, error) {
    id = request.ID
    model = request.KubernetesCluster
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
            return db.UpdateKubernetesCluster(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "kubernetes_cluster",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.KubernetesCluster.UpdateResponse{
        KubernetesCluster: model,
    }, nil
}

//RESTDeleteKubernetesCluster delete a resource using REST service.
func (service *ContrailService) RESTDeleteKubernetesCluster(c echo.Context) error {
    id := c.Param("id")
    request := &models.KubernetesClusterDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteKubernetesCluster(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteKubernetesCluster delete a resource.
func (service *ContrailService) DeleteKubernetesCluster(ctx context.Context, request *models.KubernetesClusterDeleteRequest) (*models.KubernetesClusterDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteKubernetesCluster(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.KubernetesClusterDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowKubernetesCluster a REST Show request.
func (service *ContrailService) RESTShowKubernetesCluster(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.KubernetesCluster
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListKubernetesCluster(tx, &common.ListSpec{
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
        "kubernetes_cluster": result,
    })
}

//RESTListKubernetesCluster handles a List REST service Request.
func (service *ContrailService) RESTListKubernetesCluster(c echo.Context) (error) {
    var result []*models.KubernetesCluster
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListKubernetesCluster(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "kubernetes-clusters": result,
    })
}