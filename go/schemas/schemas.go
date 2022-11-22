package schemas

type SingleResourceUri struct {
	ID uint `uri:"id" binding:"required"`
}
