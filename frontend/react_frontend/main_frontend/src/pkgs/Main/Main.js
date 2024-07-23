import "./assets/css/Main.css"
import { MainLoginForm, MainProfile, MainGoSignUp } from "./components"

export const Main = () => {
  return (
    <div className = "mainContainer">

      {/* 메인 프로필에 해당하는 컴퍼넌트 */}
      <MainProfile />

      {/* 로그인을 할 수 있는 등 메인에서 가장 많이 사용되는 컴퍼넌트 */}
      <MainLoginForm />

      {/* 회원가입을 하는 창으로 가게 하는 컴퍼넌트 */}
      <MainGoSignUp />
    </div>
  )
}