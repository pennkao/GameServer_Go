package player

// 玩家对象
type Player struct {
	// 玩家Id
	id string

	// 玩家名称
	name string

	// 玩家对应的客户端Id
	clientId int32

	// 合作商Id
	partnerId int

	// 服务器Id
	serverId int

	// 游戏版本号
	gameVersionId int

	// 资源版本号
	resourceVersionId int
}

func New(id, name string, clientId int32, partnerId, serverId, gameVersionId, resourceVersionId int) *Player {
	return &Player{
		id:                id,
		name:              name,
		clientId:          clientId,
		partnerId:         partnerId,
		serverId:          serverId,
		gameVersionId:     gameVersionId,
		resourceVersionId: resourceVersionId,
	}
}

func (p Player) Id() string {
	return p.id
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) SetName(name string) {
	p.name = name
}

func (p Player) ClientId() int32 {
	return p.clientId
}

func (p *Player) SetClientId(clientId int32) {
	p.clientId = clientId
}

func (p *Player) PartnerId() int {
	return p.partnerId
}

func (p *Player) ServerId() int {
	return p.serverId
}

func (p *Player) GameVersionId() int {
	return p.gameVersionId
}

func (p *Player) ResourceVersionId() int {
	return p.resourceVersionId
}
