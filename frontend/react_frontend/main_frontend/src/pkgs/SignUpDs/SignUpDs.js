import "./assets/css/SignUpDs.css"

import { SignUpDsHeader, SignUpDsSignUpForm } from "./components"

export const SignUpDs = () => {

  return (
    <div className = "SignUpDsContainer">
      
      {/* 회원가입 창에 머리가 되는 장소 */}
      <SignUpDsHeader />

      {/* 실질적 회원가입이 이루어 지는 장소 */}
      <SignUpDsSignUpForm />

    </div>
  )
}