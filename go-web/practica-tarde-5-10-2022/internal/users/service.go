package users

type Service interface{
	CrearUsuario(nombre string, apellido string, correo string, edad int, activo bool, fechaCreacion string)(User, error)
	GetAll() ([]User)
	GetById(id int) (User, error)
	Update(id int, nombre string, apellido string, correo string, edad int, activo bool, fechaCreacion string)(User, error)
	Delete(id int)error
	Patch(id int, nombre string, apellido string)(User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
	// va a retornar un servicio que tiene como dependencia un repositorio
}

func (s *service) CrearUsuario(
	nombre string,
	apellido string,
	correo string,
	edad int,
	activo bool,
	fechaCreacion string,
)(User, error) {
	lastId := s.repository.LastId()
	lastId ++
	return s.repository.CrearUsuario(
		lastId,
		nombre,
		apellido,
		correo,
		edad,
		activo,
		fechaCreacion,
	)
}

func (s *service) GetAll() ([]User) {
	return s.repository.GetAll()
}

func (s *service) GetById(id int) (User, error) {
	return s.repository.GetById(id)
}

func (s *service) Update(
	id int,
	nombre string,
	apellido string,
	correo string,
	edad int,
	activo bool,
	fechaCreacion string,
)(User, error){
	return s.repository.Update(
		id,
		nombre,
		apellido,
		correo,
		edad,
		activo,
		fechaCreacion,
	)
}

func (s *service) Patch(id int, nombre string, apellido string)(User, error) {
	return s.repository.Patch(id, nombre, apellido)
}

func (s *service) Delete(id int) error{
	return s.repository.Delete(id)
}






