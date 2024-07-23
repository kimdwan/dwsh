import { useMainGoSignUpHook } from "../hooks"


export const MainGoSignUp = () => {
  useMainGoSignUpHook()

  return (
    <div className = "mainGoSignUpContainer">
      <button className = "mainGoSignUpBtn">
        회원가입
      </button>
    </div>
  )
}