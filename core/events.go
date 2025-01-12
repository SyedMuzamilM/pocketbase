package core

import (
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/models/settings"
	"github.com/pocketbase/pocketbase/tools/mailer"
	"github.com/pocketbase/pocketbase/tools/search"
	"github.com/pocketbase/pocketbase/tools/subscriptions"

	"github.com/labstack/echo/v5"
)

// -------------------------------------------------------------------
// Serve events data
// -------------------------------------------------------------------

type BootstrapEvent struct {
	App App
}

type ServeEvent struct {
	App    App
	Router *echo.Echo
}

type ApiErrorEvent struct {
	HttpContext echo.Context
	Error       error
}

// -------------------------------------------------------------------
// Model DAO events data
// -------------------------------------------------------------------

type ModelEvent struct {
	Dao   *daos.Dao
	Model models.Model
}

// -------------------------------------------------------------------
// Mailer events data
// -------------------------------------------------------------------

type MailerRecordEvent struct {
	MailClient mailer.Mailer
	Message    *mailer.Message
	Record     *models.Record
	Meta       map[string]any
}

type MailerAdminEvent struct {
	MailClient mailer.Mailer
	Message    *mailer.Message
	Admin      *models.Admin
	Meta       map[string]any
}

// -------------------------------------------------------------------
// Realtime API events data
// -------------------------------------------------------------------

type RealtimeConnectEvent struct {
	HttpContext echo.Context
	Client      subscriptions.Client
}

type RealtimeDisconnectEvent struct {
	HttpContext echo.Context
	Client      subscriptions.Client
}

type RealtimeMessageEvent struct {
	HttpContext echo.Context
	Client      subscriptions.Client
	Message     *subscriptions.Message
}

type RealtimeSubscribeEvent struct {
	HttpContext   echo.Context
	Client        subscriptions.Client
	Subscriptions []string
}

// -------------------------------------------------------------------
// Settings API events data
// -------------------------------------------------------------------

type SettingsListEvent struct {
	HttpContext      echo.Context
	RedactedSettings *settings.Settings
}

type SettingsUpdateEvent struct {
	HttpContext echo.Context
	OldSettings *settings.Settings
	NewSettings *settings.Settings
}

// -------------------------------------------------------------------
// Record CRUD API events data
// -------------------------------------------------------------------

type RecordsListEvent struct {
	HttpContext echo.Context
	Collection  *models.Collection
	Records     []*models.Record
	Result      *search.Result
}

type RecordViewEvent struct {
	HttpContext echo.Context
	Record      *models.Record
}

type RecordCreateEvent struct {
	HttpContext echo.Context
	Record      *models.Record
}

type RecordUpdateEvent struct {
	HttpContext echo.Context
	Record      *models.Record
}

type RecordDeleteEvent struct {
	HttpContext echo.Context
	Record      *models.Record
}

// -------------------------------------------------------------------
// Auth Record API events data
// -------------------------------------------------------------------

type RecordAuthEvent struct {
	HttpContext echo.Context
	Record      *models.Record
	Token       string
	Meta        any
}

type RecordRequestPasswordResetEvent struct {
	HttpContext echo.Context
	Record      *models.Record
}

type RecordConfirmPasswordResetEvent struct {
	HttpContext echo.Context
	Record      *models.Record
}

type RecordRequestVerificationEvent struct {
	HttpContext echo.Context
	Record      *models.Record
}

type RecordConfirmVerificationEvent struct {
	HttpContext echo.Context
	Record      *models.Record
}

type RecordRequestEmailChangeEvent struct {
	HttpContext echo.Context
	Record      *models.Record
}

type RecordConfirmEmailChangeEvent struct {
	HttpContext echo.Context
	Record      *models.Record
}

type RecordListExternalAuthsEvent struct {
	HttpContext   echo.Context
	Record        *models.Record
	ExternalAuths []*models.ExternalAuth
}

type RecordUnlinkExternalAuthEvent struct {
	HttpContext  echo.Context
	Record       *models.Record
	ExternalAuth *models.ExternalAuth
}

// -------------------------------------------------------------------
// Admin API events data
// -------------------------------------------------------------------

type AdminsListEvent struct {
	HttpContext echo.Context
	Admins      []*models.Admin
	Result      *search.Result
}

type AdminViewEvent struct {
	HttpContext echo.Context
	Admin       *models.Admin
}

type AdminCreateEvent struct {
	HttpContext echo.Context
	Admin       *models.Admin
}

type AdminUpdateEvent struct {
	HttpContext echo.Context
	Admin       *models.Admin
}

type AdminDeleteEvent struct {
	HttpContext echo.Context
	Admin       *models.Admin
}

type AdminAuthEvent struct {
	HttpContext echo.Context
	Admin       *models.Admin
	Token       string
}

// -------------------------------------------------------------------
// Collection API events data
// -------------------------------------------------------------------

type CollectionsListEvent struct {
	HttpContext echo.Context
	Collections []*models.Collection
	Result      *search.Result
}

type CollectionViewEvent struct {
	HttpContext echo.Context
	Collection  *models.Collection
}

type CollectionCreateEvent struct {
	HttpContext echo.Context
	Collection  *models.Collection
}

type CollectionUpdateEvent struct {
	HttpContext echo.Context
	Collection  *models.Collection
}

type CollectionDeleteEvent struct {
	HttpContext echo.Context
	Collection  *models.Collection
}

type CollectionsImportEvent struct {
	HttpContext echo.Context
	Collections []*models.Collection
}

// -------------------------------------------------------------------
// Projects API events data
// -------------------------------------------------------------------
type ProjectsListEvent struct {
	HttpContext echo.Context
	Project []*models.Project
	Result      *search.Result
}

type ProjectCreateEvent struct {
	HttpContext echo.Context
	Project  *models.Project
}

// -------------------------------------------------------------------
// File API events data
// -------------------------------------------------------------------

type FileDownloadEvent struct {
	HttpContext echo.Context
	Collection  *models.Collection
	Record      *models.Record
	FileField   *schema.SchemaField
	ServedPath  string
	ServedName  string
}
