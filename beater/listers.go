// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2018-04-06 15:45:05.042968302 -0400 EDT m=+0.001733322

package beater

import (
    "github.com/google/go-github/github"
	"github.com/elastic/beats/libbeat/common"
)

func (rc *repositoryClient) ListBranches(max int) ([]*github.Branch, error) {
	opt := &github.ListOptions{}

	var results []*github.Branch

	for {
		list, resp, err := rc.client.Repositories.ListBranches(rc.ctx, rc.GetOwner(), rc.GetName(), opt)
		if err != nil {
			return results, err
		}
		
		results = append(results, list...)
		if resp.NextPage == 0 || (len(results) >= max && max > 0) {
			break
		}
		
		opt.Page = resp.NextPage
	}

	return results, nil
}

func (rc *repositoryClient) ListForks(max int) ([]*github.Repository, error) {
	opt := &github.RepositoryListForksOptions{}

	var results []*github.Repository

	for {
		list, resp, err := rc.client.Repositories.ListForks(rc.ctx, rc.GetOwner(), rc.GetName(), opt)
		if err != nil {
			return results, err
		}
		
		results = append(results, list...)
		if resp.NextPage == 0 || (len(results) >= max && max > 0) {
			break
		}
		
		opt.Page = resp.NextPage
	}

	return results, nil
}

func (rc *repositoryClient) ListComments(max int) ([]*github.RepositoryComment, error) {
	opt := &github.ListOptions{}

	var results []*github.RepositoryComment

	for {
		list, resp, err := rc.client.Repositories.ListComments(rc.ctx, rc.GetOwner(), rc.GetName(), opt)
		if err != nil {
			return results, err
		}
		
		results = append(results, list...)
		if resp.NextPage == 0 || (len(results) >= max && max > 0) {
			break
		}
		
		opt.Page = resp.NextPage
	}

	return results, nil
}

func (rc *repositoryClient) ListCommits(max int) ([]*github.RepositoryCommit, error) {
	opt := &github.CommitsListOptions{}

	var results []*github.RepositoryCommit

	for {
		list, resp, err := rc.client.Repositories.ListCommits(rc.ctx, rc.GetOwner(), rc.GetName(), opt)
		if err != nil {
			return results, err
		}
		
		results = append(results, list...)
		if resp.NextPage == 0 || (len(results) >= max && max > 0) {
			break
		}
		
		opt.Page = resp.NextPage
	}

	return results, nil
}

func (rc *repositoryClient) ListContributors(max int) ([]*github.Contributor, error) {
	opt := &github.ListContributorsOptions{}

	var results []*github.Contributor

	for {
		list, resp, err := rc.client.Repositories.ListContributors(rc.ctx, rc.GetOwner(), rc.GetName(), opt)
		if err != nil {
			return results, err
		}
		
		results = append(results, list...)
		if resp.NextPage == 0 || (len(results) >= max && max > 0) {
			break
		}
		
		opt.Page = resp.NextPage
	}

	return results, nil
}

func (rc *repositoryClient) ListDeployments(max int) ([]*github.Deployment, error) {
	opt := &github.DeploymentsListOptions{}

	var results []*github.Deployment

	for {
		list, resp, err := rc.client.Repositories.ListDeployments(rc.ctx, rc.GetOwner(), rc.GetName(), opt)
		if err != nil {
			return results, err
		}
		
		results = append(results, list...)
		if resp.NextPage == 0 || (len(results) >= max && max > 0) {
			break
		}
		
		opt.Page = resp.NextPage
	}

	return results, nil
}

func (rc *repositoryClient) ListHooks(max int) ([]*github.Hook, error) {
	opt := &github.ListOptions{}

	var results []*github.Hook

	for {
		list, resp, err := rc.client.Repositories.ListHooks(rc.ctx, rc.GetOwner(), rc.GetName(), opt)
		if err != nil {
			return results, err
		}
		
		results = append(results, list...)
		if resp.NextPage == 0 || (len(results) >= max && max > 0) {
			break
		}
		
		opt.Page = resp.NextPage
	}

	return results, nil
}

func (rc *repositoryClient) ListInvitations(max int) ([]*github.RepositoryInvitation, error) {
	opt := &github.ListOptions{}

	var results []*github.RepositoryInvitation

	for {
		list, resp, err := rc.client.Repositories.ListInvitations(rc.ctx, rc.GetOwner(), rc.GetName(), opt)
		if err != nil {
			return results, err
		}
		
		results = append(results, list...)
		if resp.NextPage == 0 || (len(results) >= max && max > 0) {
			break
		}
		
		opt.Page = resp.NextPage
	}

	return results, nil
}

func (rc *repositoryClient) ListKeys(max int) ([]*github.Key, error) {
	opt := &github.ListOptions{}

	var results []*github.Key

	for {
		list, resp, err := rc.client.Repositories.ListKeys(rc.ctx, rc.GetOwner(), rc.GetName(), opt)
		if err != nil {
			return results, err
		}
		
		results = append(results, list...)
		if resp.NextPage == 0 || (len(results) >= max && max > 0) {
			break
		}
		
		opt.Page = resp.NextPage
	}

	return results, nil
}

func (rc *repositoryClient) ListPagesBuilds(max int) ([]*github.PagesBuild, error) {
	opt := &github.ListOptions{}

	var results []*github.PagesBuild

	for {
		list, resp, err := rc.client.Repositories.ListPagesBuilds(rc.ctx, rc.GetOwner(), rc.GetName(), opt)
		if err != nil {
			return results, err
		}
		
		results = append(results, list...)
		if resp.NextPage == 0 || (len(results) >= max && max > 0) {
			break
		}
		
		opt.Page = resp.NextPage
	}

	return results, nil
}

func (rc *repositoryClient) ListProjects(max int) ([]*github.Project, error) {
	opt := &github.ProjectListOptions{}

	var results []*github.Project

	for {
		list, resp, err := rc.client.Repositories.ListProjects(rc.ctx, rc.GetOwner(), rc.GetName(), opt)
		if err != nil {
			return results, err
		}
		
		results = append(results, list...)
		if resp.NextPage == 0 || (len(results) >= max && max > 0) {
			break
		}
		
		opt.Page = resp.NextPage
	}

	return results, nil
}

func (rc *repositoryClient) ListReleases(max int) ([]*github.RepositoryRelease, error) {
	opt := &github.ListOptions{}

	var results []*github.RepositoryRelease

	for {
		list, resp, err := rc.client.Repositories.ListReleases(rc.ctx, rc.GetOwner(), rc.GetName(), opt)
		if err != nil {
			return results, err
		}
		
		results = append(results, list...)
		if resp.NextPage == 0 || (len(results) >= max && max > 0) {
			break
		}
		
		opt.Page = resp.NextPage
	}

	return results, nil
}

func (rc *repositoryClient) ListTags(max int) ([]*github.RepositoryTag, error) {
	opt := &github.ListOptions{}

	var results []*github.RepositoryTag

	for {
		list, resp, err := rc.client.Repositories.ListTags(rc.ctx, rc.GetOwner(), rc.GetName(), opt)
		if err != nil {
			return results, err
		}
		
		results = append(results, list...)
		if resp.NextPage == 0 || (len(results) >= max && max > 0) {
			break
		}
		
		opt.Page = resp.NextPage
	}

	return results, nil
}

func (rc *repositoryClient) ListTeams(max int) ([]*github.Team, error) {
	opt := &github.ListOptions{}

	var results []*github.Team

	for {
		list, resp, err := rc.client.Repositories.ListTeams(rc.ctx, rc.GetOwner(), rc.GetName(), opt)
		if err != nil {
			return results, err
		}
		
		results = append(results, list...)
		if resp.NextPage == 0 || (len(results) >= max && max > 0) {
			break
		}
		
		opt.Page = resp.NextPage
	}

	return results, nil
}

func (rc *repositoryClient) ListCollaborators(max int) ([]*github.User, error) {
	opt := &github.ListCollaboratorsOptions{}

	var results []*github.User

	for {
		list, resp, err := rc.client.Repositories.ListCollaborators(rc.ctx, rc.GetOwner(), rc.GetName(), opt)
		if err != nil {
			return results, err
		}
		
		results = append(results, list...)
		if resp.NextPage == 0 || (len(results) >= max && max > 0) {
			break
		}
		
		opt.Page = resp.NextPage
	}

	return results, nil
}

func (rc *repositoryClient) ListParticipation() (*github.RepositoryParticipation, error) {
	result, _, err := rc.client.Repositories.ListParticipation(rc.ctx, rc.GetOwner(), rc.GetName())
	return result, err
}

func (rc *repositoryClient) ListLanguages() (map[string]int, error) {
	result, _, err := rc.client.Repositories.ListLanguages(rc.ctx, rc.GetOwner(), rc.GetName())
	return result, err
}

func (rc *repositoryClient) ListContributorsStats() ([]*github.ContributorStats, error) {
	result, _, err := rc.client.Repositories.ListContributorsStats(rc.ctx, rc.GetOwner(), rc.GetName())
	return result, err
}

func (rc *repositoryClient) ListPunchCard() ([]*github.PunchCard, error) {
	result, _, err := rc.client.Repositories.ListPunchCard(rc.ctx, rc.GetOwner(), rc.GetName())
	return result, err
}

func (rc *repositoryClient) ListTrafficPaths() ([]*github.TrafficPath, error) {
	result, _, err := rc.client.Repositories.ListTrafficPaths(rc.ctx, rc.GetOwner(), rc.GetName())
	return result, err
}

func (rc *repositoryClient) ListTrafficReferrers() ([]*github.TrafficReferrer, error) {
	result, _, err := rc.client.Repositories.ListTrafficReferrers(rc.ctx, rc.GetOwner(), rc.GetName())
	return result, err
}

func (rc *repositoryClient) ListAllTopics() ([]string, error) {
	result, _, err := rc.client.Repositories.ListAllTopics(rc.ctx, rc.GetOwner(), rc.GetName())
	return result, err
}

func (rc *repositoryClient) ListCodeFrequency() ([]*github.WeeklyStats, error) {
	result, _, err := rc.client.Repositories.ListCodeFrequency(rc.ctx, rc.GetOwner(), rc.GetName())
	return result, err
}

func (rc *repositoryClient) ListCommitActivity() ([]*github.WeeklyCommitActivity, error) {
	result, _, err := rc.client.Repositories.ListCommitActivity(rc.ctx, rc.GetOwner(), rc.GetName())
	return result, err
}

func (bt *Githubbeat) collectDownloads(rc *repositoryClient) common.MapStr {
	if ! bt.config.Downloads.Enabled {
		return nil
	}

	rawResults, err := rc.ListReleases(bt.config.Downloads.Max)

	formatted := bt.extractDownloads(rawResults, err)
	
	return appendError(formatted, err)
	
}

func (bt *Githubbeat) collectParticipation(rc *repositoryClient) common.MapStr {
	if ! bt.config.Participation.Enabled {
		return nil
	}

	rawResults, err := rc.ListParticipation()

	formatted := bt.extractParticipation(rawResults, err)
	
	return appendError(formatted, err)
	
}

func (bt *Githubbeat) collectForks(rc *repositoryClient) common.MapStr {
	if ! bt.config.Forks.Enabled {
		return nil
	}

	rawResults, err := rc.ListForks(bt.config.Forks.Max)

	formatted := bt.extractForks(rawResults, err)
	
	return 	createListMapStr(formatted, err, bt.config.Forks.List)
	
}

func (bt *Githubbeat) collectLanguages(rc *repositoryClient) common.MapStr {
	if ! bt.config.Languages.Enabled {
		return nil
	}

	rawResults, err := rc.ListLanguages()

	formatted := bt.extractLanguages(rawResults, err)
	
	return 	createListMapStr(formatted, err, bt.config.Languages.List)
	
}

func (bt *Githubbeat) collectContributors(rc *repositoryClient) common.MapStr {
	if ! bt.config.Contributors.Enabled {
		return nil
	}

	rawResults, err := rc.ListContributors(bt.config.Contributors.Max)

	formatted := bt.extractContributors(rawResults, err)
	
	return 	createListMapStr(formatted, err, bt.config.Contributors.List)
	
}

func (bt *Githubbeat) collectBranches(rc *repositoryClient) common.MapStr {
	if ! bt.config.Branches.Enabled {
		return nil
	}

	rawResults, err := rc.ListBranches(bt.config.Branches.Max)

	formatted := bt.extractBranches(rawResults, err)
	
	return 	createListMapStr(formatted, err, bt.config.Branches.List)
	
}
