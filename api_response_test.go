/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package bitbucketv1

import (
	"net/http"
	"reflect"
	"testing"
)

func TestSSHKey_String(t *testing.T) {
	type fields struct {
		ID    int
		Text  string
		Label string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &SSHKey{
				ID:    tt.fields.ID,
				Text:  tt.fields.Text,
				Label: tt.fields.Label,
			}
			if got := k.String(); got != tt.want {
				t.Errorf("SSHKey.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCommitsResponse(t *testing.T) {
	type args struct {
		r *APIResponse
	}
	tests := []struct {
		name    string
		args    args
		want    []Commit
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCommitsResponse(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCommitsResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCommitsResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTagsResponse(t *testing.T) {
	type args struct {
		r *APIResponse
	}
	tests := []struct {
		name    string
		args    args
		want    []Tag
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTagsResponse(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTagsResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTagsResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBranchesResponse(t *testing.T) {
	type args struct {
		r *APIResponse
	}
	tests := []struct {
		name    string
		args    args
		want    []Branch
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBranchesResponse(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBranchesResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBranchesResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRepositoriesResponse(t *testing.T) {
	type args struct {
		r *APIResponse
	}
	tests := []struct {
		name    string
		args    args
		want    []Repository
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRepositoriesResponse(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRepositoriesResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRepositoriesResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRepositoryResponse(t *testing.T) {
	type args struct {
		r *APIResponse
	}
	tests := []struct {
		name    string
		args    args
		want    Repository
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRepositoryResponse(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRepositoryResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRepositoryResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDiffResponse(t *testing.T) {
	type args struct {
		r *APIResponse
	}
	tests := []struct {
		name    string
		args    args
		want    Diff
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDiffResponse(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDiffResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDiffResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSSHKeysResponse(t *testing.T) {
	type args struct {
		r *APIResponse
	}
	tests := []struct {
		name    string
		args    args
		want    []SSHKey
		wantErr bool
	}{
		{
			name: "Empty list",
			args: args{
				r: &APIResponse{
					Values: map[string]interface{}{"values": []interface{}{}},
				},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Single sshkey",
			args: args{
				r: &APIResponse{
					Values: map[string]interface{}{
						"values": []interface{}{map[string]interface{}{
							"id":    1,
							"text":  "ssh ....",
							"label": "My-ssh-key",
						}},
					},
				},
			},
			want: []SSHKey{
				SSHKey{
					ID:    1,
					Text:  "ssh ....",
					Label: "My-ssh-key",
				},
			},
			wantErr: false,
		},
		{
			name: "Bad response",
			args: args{
				r: &APIResponse{
					Values: map[string]interface{}{"values": "not an array"},
				},
			},
			want:    nil,
			wantErr: true,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSSHKeysResponse(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSSHKeysResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSSHKeysResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPullRequestResponse(t *testing.T) {
	type args struct {
		r *APIResponse
	}
	tests := []struct {
		name    string
		args    args
		want    PullRequest
		wantErr bool
	}{
		{
			name: "Empty list",
			args: args{
				r: &APIResponse{
					Values: map[string]interface{}{},
				},
			},
			want:    PullRequest{},
			wantErr: false,
		},
		{
			name: "Simple PullRequest",
			args: args{
				r: &APIResponse{
					Values: map[string]interface{}{
						"id": 1,
					},
				},
			},
			want: PullRequest{
				ID: 1,
			},
			wantErr: false,
		},
		{
			name: "Good response",
			args: args{
				r: &APIResponse{
					Values: map[string]interface{}{
						"id":          101,
						"version":     1,
						"title":       "Talking Nerdy",
						"description": "It’s a kludge, but put the tuple from the database in the cache.",
						"state":       "OPEN",
						"open":        true,
						"closed":      false,
						"createdDate": 1359075920,
						"updatedDate": 1359085920,
						"fromRef": map[string]interface{}{
							"id": "refs/heads/feature-ABC-123",
							"repository": map[string]interface{}{
								"slug": "my-repo",
								"name": nil,
								"project": map[string]interface{}{
									"key": "PRJ",
								},
							},
						},
						"toRef": map[string]interface{}{
							"id": "refs/heads/master",
							"repository": map[string]interface{}{
								"slug": "my-repo",
								"name": nil,
								"project": map[string]interface{}{
									"key": "PRJ",
								},
							},
						},
						"locked": false,
						"author": map[string]interface{}{
							"user": map[string]interface{}{
								"name":         "tom",
								"emailAddress": "tom@example.com",
								"id":           115026,
								"displayName":  "Tom",
								"active":       true,
								"slug":         "tom",
								"type":         "NORMAL",
							},
							"role":     "AUTHOR",
							"approved": true,
							"status":   "APPROVED",
						},
						"reviewers": []map[string]interface{}{
							{
								"user": map[string]interface{}{
									"name":         "jcitizen",
									"emailAddress": "jane@example.com",
									"id":           101,
									"displayName":  "Jane Citizen",
									"active":       true,
									"slug":         "jcitizen",
									"type":         "NORMAL",
								},
								"lastReviewedCommit": "7549846524f8aed2bd1c0249993ae1bf9d3c9998",
								"role":               "REVIEWER",
								"approved":           true,
								"status":             "APPROVED",
							},
						},
						"participants": []map[string]interface{}{
							{
								"user": map[string]interface{}{
									"name":         "dick",
									"emailAddress": "dick@example.com",
									"id":           3083181,
									"displayName":  "Dick",
									"active":       true,
									"slug":         "dick",
									"type":         "NORMAL",
								},
								"role":     "PARTICIPANT",
								"approved": false,
								"status":   "UNAPPROVED",
							},
							{
								"user": map[string]interface{}{
									"name":         "harry",
									"emailAddress": "harry@example.com",
									"id":           99049120,
									"displayName":  "Harry",
									"active":       true,
									"slug":         "harry",
									"type":         "NORMAL",
								},
								"role":     "PARTICIPANT",
								"approved": true,
								"status":   "APPROVED",
							},
						},
						"links": map[string]interface{}{
							"self": []map[string]interface{}{
								{
									"href": "http://link/to/pullrequest",
								},
							},
						},
					},
				},
			},
			want: PullRequest{
				ID:          101,
				Version:     1,
				Title:       "Talking Nerdy",
				Description: "It’s a kludge, but put the tuple from the database in the cache.",
				State:       "OPEN",
				Open:        true,
				Closed:      false,
				CreatedDate: 1359075920,
				UpdatedDate: 1359085920,
				FromRef: PullRequestRef{
					ID: "refs/heads/feature-ABC-123",
					Repository: Repository{
						Slug: "my-repo",
						Project: &Project{
							Key: "PRJ",
						},
					},
				},
				ToRef: PullRequestRef{
					ID: "refs/heads/master",
					Repository: Repository{
						Slug: "my-repo",
						Project: &Project{
							Key: "PRJ",
						},
					},
				},
				Locked: false,
				Author: &UserWithMetadata{
					User: UserWithLinks{
						Name:         "tom",
						EmailAddress: "tom@example.com",
						ID:           115026,
						DisplayName:  "Tom",
						Active:       true,
						Slug:         "tom",
						Type:         "NORMAL",
					},
					Role:     "AUTHOR",
					Approved: true,
					Status:   "APPROVED",
				},
				Reviewers: []UserWithMetadata{
					{
						User: UserWithLinks{
							Name:         "jcitizen",
							EmailAddress: "jane@example.com",
							ID:           101,
							DisplayName:  "Jane Citizen",
							Active:       true,
							Slug:         "jcitizen",
							Type:         "NORMAL",
						},
						Role:               "REVIEWER",
						Approved:           true,
						Status:             "APPROVED",
						LastReviewedCommit: "7549846524f8aed2bd1c0249993ae1bf9d3c9998",
					},
				},
				Participants: []UserWithMetadata{
					{
						User: UserWithLinks{
							Name:         "dick",
							EmailAddress: "dick@example.com",
							ID:           3083181,
							DisplayName:  "Dick",
							Active:       true,
							Slug:         "dick",
							Type:         "NORMAL",
						},
						Role:     "PARTICIPANT",
						Approved: false,
						Status:   "UNAPPROVED",
					},
					{
						User: UserWithLinks{
							Name:         "harry",
							EmailAddress: "harry@example.com",
							ID:           99049120,
							DisplayName:  "Harry",
							Active:       true,
							Slug:         "harry",
							Type:         "NORMAL",
						},
						Role:     "PARTICIPANT",
						Approved: true,
						Status:   "APPROVED",
					},
				},
				Links: Links{
					Self: []SelfLink{
						{
							Href: "http://link/to/pullrequest",
						},
					},
				},
			},
			wantErr: false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPullRequestResponse(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPullRequestResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPullRequestResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPullRequestsResponse(t *testing.T) {
	type args struct {
		r *APIResponse
	}
	tests := []struct {
		name    string
		args    args
		want    []PullRequest
		wantErr bool
	}{
		{
			name: "Empty list",
			args: args{
				r: &APIResponse{
					Values: map[string]interface{}{"values": []interface{}{}},
				},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Simple PullRequest",
			args: args{
				r: &APIResponse{
					Values: map[string]interface{}{
						"values": []interface{}{map[string]interface{}{
							"id": 1,
						}},
					},
				},
			},
			want: []PullRequest{{
				ID: 1,
			},
			},
			wantErr: false,
		},
		{
			name: "Bad response",
			args: args{
				r: &APIResponse{
					Values: map[string]interface{}{"values": "not an array"},
				},
			},
			want:    nil,
			wantErr: true,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPullRequestsResponse(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPullRequestsResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPullRequestsResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWebhooksResponse(t *testing.T) {
	type args struct {
		r *APIResponse
	}
	tests := []struct {
		name    string
		args    args
		want    []Webhook
		wantErr bool
	}{
		{
			name: "Empty list",
			args: args{
				r: &APIResponse{
					Values: map[string]interface{}{"values": []interface{}{}},
				},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Single webhook",
			args: args{
				r: &APIResponse{
					Values: map[string]interface{}{
						"values": []interface{}{map[string]interface{}{
							"id":     1,
							"name":   "foo",
							"url":    "http://bitbucket.localhost/hook",
							"active": false,
							"events": []string{"repo:modified"},
							"configuration": map[string]interface{}{
								"secret": "password",
							},
						}},
					},
				},
			},
			want: []Webhook{
				Webhook{
					ID:     1,
					Name:   "foo",
					Url:    "http://bitbucket.localhost/hook",
					Active: false,
					Events: []string{"repo:modified"},
					Configuration: WebhookConfiguration{
						Secret: "password",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Bad response",
			args: args{
				r: &APIResponse{
					Values: map[string]interface{}{"values": "not an array"},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetWebhooksResponse(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWebhooksResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWebhooksResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUsersPermissionResponse(t *testing.T) {
	type args struct {
		r *APIResponse
	}
	tests := []struct {
		name    string
		args    args
		want    []UserPermission
		wantErr bool
	}{
		{
			name: "Empty list",
			args: args{
				r: &APIResponse{
					Values: map[string]interface{}{"values": []interface{}{}},
				},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Bad response",
			args: args{
				r: &APIResponse{
					Values: map[string]interface{}{"values": "not an array"},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Single user permission",
			args: args{
				r: &APIResponse{
					Values: map[string]interface{}{
						"values": []interface{}{map[string]interface{}{
							"user": map[string]interface{}{
								"name":         "jcitizen",
								"emailAddress": "jane@example.com",
								"id":           101,
								"displayName":  "Jane Citizen",
								"active":       true,
								"slug":         "jcitizen",
								"type":         "NORMAL",
							},
							"permission": "REPO_ADMIN",
						}},
					},
				},
			},
			want: []UserPermission{
				UserPermission{
					User: User{
						Name:         "jcitizen",
						EmailAddress: "jane@example.com",
						ID:           101,
						DisplayName:  "Jane Citizen",
						Active:       true,
						Slug:         "jcitizen",
						Type:         "NORMAL",
					},
					Permission: PermissionRepositoryAdmin.String(),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUsersPermissionResponse(tt.args.r)
			if err != nil && !tt.wantErr {
				t.Errorf("GetUsersPermissionResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsersPermissionResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUsersResponse(t *testing.T) {
	type args struct {
		r *APIResponse
	}
	tests := []struct {
		name    string
		args    args
		want    []User
		wantErr bool
	}{
		{
			name: "Empty list",
			args: args{
				r: &APIResponse{
					Values: map[string]interface{}{"values": []interface{}{}},
				},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Bad response",
			args: args{
				r: &APIResponse{
					Values: map[string]interface{}{"values": "not an array"},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Single group permission",
			args: args{
				r: &APIResponse{
					Values: map[string]interface{}{
						"values": []interface{}{
							map[string]interface{}{
								"name":                        "jcitizen",
								"emailAddress":                "jane@example.com",
								"id":                          101,
								"displayName":                 "Jane Citizen",
								"active":                      true,
								"slug":                        "jcitizen",
								"type":                        "NORMAL",
								"directoryName":               "Bitbucket Internal Directory",
								"deletable":                   true,
								"lastAuthenticationTimestamp": 1368145580548,
								"mutableDetails":              true,
								"mutableGroups":               true,
							},
						},
					},
				},
			},
			want: []User{
				User{
					Name:                        "jcitizen",
					EmailAddress:                "jane@example.com",
					ID:                          101,
					DisplayName:                 "Jane Citizen",
					Active:                      true,
					Slug:                        "jcitizen",
					Type:                        "NORMAL",
					DirectoryName:               "Bitbucket Internal Directory",
					Deletable:                   true,
					LastAuthenticationTimestamp: 1368145580548,
					MutableDetails:              true,
					MutableGroups:               true,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUsersResponse(tt.args.r)
			if err != nil && !tt.wantErr {
				t.Errorf("GetGroupsPermissionResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGroupsPermissionResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetGroupsPermissionResponse(t *testing.T) {
	type args struct {
		r *APIResponse
	}
	tests := []struct {
		name    string
		args    args
		want    []GroupPermission
		wantErr bool
	}{
		{
			name: "Empty list",
			args: args{
				r: &APIResponse{
					Values: map[string]interface{}{"values": []interface{}{}},
				},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Bad response",
			args: args{
				r: &APIResponse{
					Values: map[string]interface{}{"values": "not an array"},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Single group permission",
			args: args{
				r: &APIResponse{
					Values: map[string]interface{}{
						"values": []interface{}{map[string]interface{}{
							"group": map[string]interface{}{
								"name": "group1",
							},
							"permission": "REPO_ADMIN",
						}},
					},
				},
			},
			want: []GroupPermission{
				GroupPermission{
					Group: Group{
						Name: "group1",
					},
					Permission: PermissionRepositoryAdmin.String(),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetGroupsPermissionResponse(tt.args.r)
			if err != nil && !tt.wantErr {
				t.Errorf("GetGroupsPermissionResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGroupsPermissionResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAPIResponse(t *testing.T) {
	type args struct {
		r *http.Response
	}
	tests := []struct {
		name string
		args args
		want *APIResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAPIResponse(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAPIResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAPIResponseWithError(t *testing.T) {
	type args struct {
		r   *http.Response
		b   []byte
		err error
	}
	tests := []struct {
		name    string
		args    args
		want    *APIResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAPIResponseWithError(tt.args.r, tt.args.b, tt.args.err)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAPIResponseWithError() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAPIResponseWithError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBitbucketAPIResponse(t *testing.T) {
	type args struct {
		r *http.Response
	}
	tests := []struct {
		name    string
		args    args
		want    *APIResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBitbucketAPIResponse(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBitbucketAPIResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBitbucketAPIResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasNextPage(t *testing.T) {
	type args struct {
		r *APIResponse
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want2 int
	}{
		{
			name: "Empty list",
			args: args{
				r: &APIResponse{
					Values: map[string]interface{}{},
				},
			},
			want:  false,
			want2: 0,
		},
		{
			name: "Bad response",
			args: args{
				r: &APIResponse{
					Values: map[string]interface{}{"values": "not an array"},
				},
			},
			want:  false,
			want2: 0,
		},
		{
			name: "Last Page",
			args: args{
				r: &APIResponse{
					Values: map[string]interface{}{
						"size":       1,
						"limit":      25,
						"isLastPage": true,
						"values":     []interface{}{},
					},
				},
			},
			want:  false,
			want2: 0,
		},
		{
			name: "Hast Last Page",
			args: args{
				r: &APIResponse{
					Values: map[string]interface{}{
						"size":          1,
						"limit":         25,
						"isLastPage":    false,
						"nextPageStart": float64(1203),
						"values":        []interface{}{},
					},
				},
			},
			want:  true,
			want2: 1203,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HasNextPage(tt.args.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBitbucketAPIResponse() isLastPage = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(err, tt.want2) {
				t.Errorf("NewBitbucketAPIResponse() nextPageStart = %v, want %v", err, tt.want2)
			}
		})
	}
}
