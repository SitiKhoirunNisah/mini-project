package view

type UserView struct {
    // Implementasi sesuai kebutuhan
}

func (v *UserView) ShowRegistrationSuccessMessage() {
    // Implementasi tampilan pesan sukses registrasi
}

func (v *UserView) ShowRegistrationError(message string) {
    // Implementasi tampilan pesan kesalahan registrasi
}

func (v *UserView) ShowLoginSuccessMessage(username string) {
    // Implementasi tampilan pesan sukses login
}

func (v *UserView) ShowLoginError(message string) {
    // Implementasi tampilan pesan kesalahan login
}

func (v *UserView) ShowLogoutMessage() {
    // Implementasi tampilan pesan logout
}
