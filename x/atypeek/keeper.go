package atypeek

import (
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/theuncharted/atypeek_blockchain/x/atypeek/tools"
)

// Keeper manages transfers between accounts
type Keeper struct {
	cdc                 *codec.Codec
	profileStoreKey     sdk.StoreKey
	projectStoreKey     sdk.StoreKey
	skillStoreKey       sdk.StoreKey
	courseStoreKey      sdk.StoreKey
	endorsementStoreKey sdk.StoreKey
}

// NewKeeper returns a new Keeper
func NewKeeper(cdc *codec.Codec, profileStoreKey sdk.StoreKey, projectStoreKey sdk.StoreKey, skillStoreKey sdk.StoreKey, courseStoreKey sdk.StoreKey, endorsementStoreKey sdk.StoreKey) Keeper {
	return Keeper{cdc: cdc, profileStoreKey: profileStoreKey, projectStoreKey: projectStoreKey, skillStoreKey: skillStoreKey, courseStoreKey: courseStoreKey, endorsementStoreKey: endorsementStoreKey}
}

func (k Keeper) SetProfile(ctx sdk.Context, r Profile) {
	if r.Owner.Empty() {
		return
	}
	fmt.Printf("set profile with owner %s\n", r.Owner.String())
	store := ctx.KVStore(k.profileStoreKey)
	store.Set([]byte(r.Owner.String()), k.cdc.MustMarshalBinaryBare(r))
}

func (k Keeper) GetProfile(ctx sdk.Context, owner sdk.AccAddress) Profile {
	fmt.Printf("get profile with owner %s", owner.String())
	store := ctx.KVStore(k.profileStoreKey)
	if !store.Has([]byte(owner.String())) {
		profile := NewProfile()
		profile.Owner = owner
		return profile
	}
	bz := store.Get([]byte(owner.String()))
	var profile Profile
	k.cdc.MustUnmarshalBinaryBare(bz, &profile)
	return profile
}

func (k Keeper) AddProject(ctx sdk.Context, p Project, owner sdk.AccAddress) error {
	if owner.Empty() {
		errors.New("owner not provided")
	}

	profile := k.GetProfile(ctx, owner)
	if tools.Contains(profile.Projects, p.Id) {
		return errors.New("project exists already")
	} else {
		fmt.Println("adding project to profile")
		store := ctx.KVStore(k.projectStoreKey)
		store.Set([]byte(p.Id), k.cdc.MustMarshalBinaryBare(p))
		profile.Projects = append(profile.Projects, p.Id)
		fmt.Println("adding profile to store")
		k.SetProfile(ctx, profile)
		return nil
	}
}

func (k Keeper) GetProject(ctx sdk.Context, id string) (Project, error) {
	store := ctx.KVStore(k.projectStoreKey)
	if !store.Has([]byte(id)) {
		return NewProject(), errors.New("No project found")
	}
	bz := store.Get([]byte(id))
	var project Project
	k.cdc.MustUnmarshalBinaryBare(bz, &project)
	return project, nil
}

func (k Keeper) SetProject(ctx sdk.Context, p Project) error {
	store := ctx.KVStore(k.projectStoreKey)
	store.Set([]byte(p.Id), k.cdc.MustMarshalBinaryBare(p))
	return nil
}

func (k Keeper) AddSkill(ctx sdk.Context, idProject string, s Skill, owner sdk.AccAddress) error {
	if owner.Empty() {
		errors.New("owner not provided")
	}

	project, err := k.GetProject(ctx, idProject)
	if err == nil {
		if project.Owner.Equals(owner) {
			store := ctx.KVStore(k.skillStoreKey)
			store.Set([]byte(s.Id), k.cdc.MustMarshalBinaryBare(s))
			project.Skills = append(project.Skills, s.Id)
			k.SetProject(ctx, project)
		}

		profile := k.GetProfile(ctx, owner)
		profile.Skills = append(profile.Skills, s.Id)
		k.SetProfile(ctx, profile)
	}

	return nil
}

func (k Keeper) GetSkill(ctx sdk.Context, id string) (Skill, error) {
	store := ctx.KVStore(k.skillStoreKey)
	if !store.Has([]byte(id)) {
		return NewSkill(), errors.New("No skill found")
	}
	bz := store.Get([]byte(id))
	var skill Skill
	k.cdc.MustUnmarshalBinaryBare(bz, &skill)
	return skill, nil
}

func (k Keeper) SetSkill(ctx sdk.Context, s Skill) error {
	store := ctx.KVStore(k.skillStoreKey)
	store.Set([]byte(s.Id), k.cdc.MustMarshalBinaryBare(s))
	return nil
}

func (k Keeper) AddCourse(ctx sdk.Context, idSkill string, c Course, owner sdk.AccAddress) error {
	if owner.Empty() {
		errors.New("owner not provided")
	}

	skill, err := k.GetSkill(ctx, idSkill)
	if err == nil {
		if skill.Owner.Equals(owner) {
			store := ctx.KVStore(k.courseStoreKey)
			store.Set([]byte(c.Id), k.cdc.MustMarshalBinaryBare(c))
			skill.Courses = append(skill.Courses, c.Id)
			k.SetSkill(ctx, skill)
		}

		profile := k.GetProfile(ctx, owner)
		profile.Courses = append(profile.Courses, c.Id)
		k.SetProfile(ctx, profile)
	}

	return nil
}

func (k Keeper) GetCourse(ctx sdk.Context, id string) (Course, error) {
	store := ctx.KVStore(k.courseStoreKey)
	if !store.Has([]byte(id)) {
		return NewCourse(), errors.New("No course found")
	}
	bz := store.Get([]byte(id))
	var course Course
	k.cdc.MustUnmarshalBinaryBare(bz, &course)
	return course, nil
}

func (k Keeper) AddEndorsement(ctx sdk.Context, idSkill string, e Endorsement, owner sdk.AccAddress) error {
	skill, err := k.GetSkill(ctx, idSkill)
	if err == nil {
		if skill.Owner.Equals(owner) {
			store := ctx.KVStore(k.endorsementStoreKey)
			store.Set([]byte(e.Id), k.cdc.MustMarshalBinaryBare(e))
			skill.Endorsements = append(skill.Endorsements, e.Id)
			k.SetSkill(ctx, skill)
		}

		profile := k.GetProfile(ctx, owner)
		profile.Endorsements = append(profile.Endorsements, e.Id)
		k.SetProfile(ctx, profile)
	}

	return nil
}

func (k Keeper) GetEndorsement(ctx sdk.Context, id string) (Endorsement, error) {
	store := ctx.KVStore(k.endorsementStoreKey)
	if !store.Has([]byte(id)) {
		return NewEndorsement(), errors.New("No endorsement found")
	}
	bz := store.Get([]byte(id))
	var endorsement Endorsement
	k.cdc.MustUnmarshalBinaryBare(bz, &endorsement)
	return endorsement, nil
}
