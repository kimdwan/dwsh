import { SignUpGuestHeader, SignUpGuestSignUpForm } from "./components"


export const SignUpGuest = () => {
  return (
    <div className = "SignUpGuestContainer">
      
      {/* guest용 회원가입의 헤더 */}
      <SignUpGuestHeader />

      {/* guest용 회원가입 폼 */}
      <SignUpGuestSignUpForm />

    </div>
  )
}