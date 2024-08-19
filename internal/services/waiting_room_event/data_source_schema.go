// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_event

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &WaitingRoomEventDataSource{}

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"event_id": schema.StringAttribute{
				Optional: true,
			},
			"waiting_room_id": schema.StringAttribute{
				Optional: true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"created_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"description": schema.StringAttribute{
				Description: "A note that you can use to add more details about the event.",
				Computed:    true,
			},
			"modified_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"shuffle_at_event_start": schema.BoolAttribute{
				Description: "If enabled, users in the prequeue will be shuffled randomly at the `event_start_time`. Requires that `prequeue_start_time` is not null. This is useful for situations when many users will join the event prequeue at the same time and you want to shuffle them to ensure fairness. Naturally, it makes the most sense to enable this feature when the `queueing_method` during the event respects ordering such as **fifo**, or else the shuffling may be unnecessary.",
				Computed:    true,
			},
			"suspended": schema.BoolAttribute{
				Description: "Suspends or allows an event. If set to `true`, the event is ignored and traffic will be handled based on the waiting room configuration.",
				Computed:    true,
			},
			"custom_page_html": schema.StringAttribute{
				Description: "If set, the event will override the waiting room's `custom_page_html` property while it is active. If null, the event will inherit it.",
				Computed:    true,
				Optional:    true,
			},
			"disable_session_renewal": schema.BoolAttribute{
				Description: "If set, the event will override the waiting room's `disable_session_renewal` property while it is active. If null, the event will inherit it.",
				Computed:    true,
				Optional:    true,
			},
			"event_end_time": schema.StringAttribute{
				Description: "An ISO 8601 timestamp that marks the end of the event.",
				Computed:    true,
				Optional:    true,
			},
			"event_start_time": schema.StringAttribute{
				Description: "An ISO 8601 timestamp that marks the start of the event. At this time, queued users will be processed with the event's configuration. The start time must be at least one minute before `event_end_time`.",
				Computed:    true,
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"name": schema.StringAttribute{
				Description: "A unique name to identify the event. Only alphanumeric characters, hyphens and underscores are allowed.",
				Computed:    true,
				Optional:    true,
			},
			"new_users_per_minute": schema.Int64Attribute{
				Description: "If set, the event will override the waiting room's `new_users_per_minute` property while it is active. If null, the event will inherit it. This can only be set if the event's `total_active_users` property is also set.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(200, 2147483647),
				},
			},
			"prequeue_start_time": schema.StringAttribute{
				Description: "An ISO 8601 timestamp that marks when to begin queueing all users before the event starts. The prequeue must start at least five minutes before `event_start_time`.",
				Computed:    true,
				Optional:    true,
			},
			"queueing_method": schema.StringAttribute{
				Description: "If set, the event will override the waiting room's `queueing_method` property while it is active. If null, the event will inherit it.",
				Computed:    true,
				Optional:    true,
			},
			"session_duration": schema.Int64Attribute{
				Description: "If set, the event will override the waiting room's `session_duration` property while it is active. If null, the event will inherit it.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 30),
				},
			},
			"total_active_users": schema.Int64Attribute{
				Description: "If set, the event will override the waiting room's `total_active_users` property while it is active. If null, the event will inherit it. This can only be set if the event's `new_users_per_minute` property is also set.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(200, 2147483647),
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"zone_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"waiting_room_id": schema.StringAttribute{
						Required: true,
					},
					"page": schema.StringAttribute{
						Description: "Page number of paginated results.",
						Computed:    true,
						Optional:    true,
						CustomType:  jsontypes.NormalizedType{},
					},
					"per_page": schema.StringAttribute{
						Description: "Maximum number of results per page. Must be a multiple of 5.",
						Computed:    true,
						Optional:    true,
						CustomType:  jsontypes.NormalizedType{},
					},
				},
			},
		},
	}
}

func (d *WaitingRoomEventDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *WaitingRoomEventDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
