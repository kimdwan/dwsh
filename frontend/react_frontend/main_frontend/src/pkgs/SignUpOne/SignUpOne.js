import "./assets/css/SignUpOne.css"
import { SignUpOneIntro, SignUpOneWho  } from "./components"

export const SignUpOne = () => {
  return (
    <div>
      
      {/* 회원가입 프로필이 보임 */}
      <SignUpOneIntro />

      {/* 회원가입 경로를 선택하게 해줌 */}
      <SignUpOneWho />
      
    </div>
  )
}