package response

import "github.com/mauriciomartinezc/real-estate-mc-auth/domain"

type CompanyUserResponse struct {
	ID        string         `json:"id"`
	UserID    string         `json:"user_id"`
	CompanyID string         `json:"company_id"`
	Roles     []RoleResponse `json:"roles,omitempty"`
	User      UserResponse   `json:"user,omitempty"`
	Creator   UserResponse   `json:"creator,omitempty"`
	Updater   UserResponse   `json:"updater,omitempty"`
	CreatedAt int64          `json:"created_at"`
	UpdatedAt int64          `json:"updated_at"`
}

func ToCompanyUserResponse(companyUser domain.CompanyUser) CompanyUserResponse {
	// Verificar si companyUser.Roles es nil
	var roles []RoleResponse
	if companyUser.Roles != nil {
		roles = make([]RoleResponse, len(*companyUser.Roles))
		for i, role := range *companyUser.Roles {
			roles[i] = RoleResponse{
				ID:   role.ID.String(),
				Name: role.Name,
				Slug: role.Slug,
			}
		}
	}

	// Verificar si companyUser.User y su Profile son nil
	userResponse := UserResponse{}
	if companyUser.User != nil && companyUser.User.Profile != nil {
		userResponse.Profile = ProfileResponse{
			FirstName: companyUser.User.Profile.FirstName,
			LastName:  companyUser.User.Profile.LastName,
		}
	}

	// Verificar si companyUser.Creator y su Profile son nil
	creatorResponse := UserResponse{}
	if companyUser.Creator != nil && companyUser.Creator.Profile != nil {
		creatorResponse = UserResponse{
			Profile: ProfileResponse{
				FirstName: companyUser.Creator.Profile.FirstName,
				LastName:  companyUser.Creator.Profile.LastName,
			},
		}
	}

	// Verificar si companyUser.Updater y su Profile son nil
	updaterResponse := UserResponse{}
	if companyUser.Updater != nil && companyUser.Updater.Profile != nil {
		updaterResponse = UserResponse{
			Profile: ProfileResponse{
				FirstName: companyUser.Updater.Profile.FirstName,
				LastName:  companyUser.Updater.Profile.LastName,
			},
		}
	}

	return CompanyUserResponse{
		ID:        companyUser.ID.String(),
		UserID:    companyUser.UserId,
		CompanyID: companyUser.CompanyId,
		Roles:     roles,
		User:      userResponse,
		Creator:   creatorResponse,
		Updater:   updaterResponse,
		CreatedAt: companyUser.CreatedAt,
		UpdatedAt: companyUser.UpdatedAt,
	}
}
