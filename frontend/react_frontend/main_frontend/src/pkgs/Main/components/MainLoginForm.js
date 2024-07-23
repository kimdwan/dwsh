import { useMainLoginFormHook } from "../hooks"

export const MainLoginForm = () => {
  const { register, handleSubmit, onSubmit, emailErrorMsg, passwordErrorMsg } = useMainLoginFormHook()

  return (
    <div className = "mainLoginFormContainer">

      {/* 유저에게 설명을 해주는 칸 */}
      <div className = "mainLoginFormLoginIntro">
        <h1 className = "mainLoginFormLoginIntroValue">저희 홈페이지를 방문해주셔서 감사합니다.</h1>
      </div>

      
      {/* 로그인 폼 */}
      <div className = "mainLoginFormLoginBox">
        <form onSubmit = {handleSubmit(onSubmit)}>

          {/* 이메일와 관련된 박스 */}
          <div className = "mainLoginFormEmailBox">
            <div className = "mainLoginFormEmailIntro">
              이메일: 
            </div>
            <input 
              type = "text"
              className = "mainLoginFormEmailInputBox"
              {...register("email")}
            />
            <div className = "mainLoginFormEmailErrorBox"> 
              <p className = "mainLoginFormError">{emailErrorMsg}</p>
            </div>
          </div>  
          
          {/* 비밀번호와 관련된 박스 */}
          <div className = "mainLoginFormPasswordBox">
            <div className = "mainLoginFormPasswordIntro">
              비밀번호: 
            </div>
            <input 
              type = "password"
              className = "mainLoginFormPasswordInputBox"
              { ...register("password") }
            />
            <div className = "mainLoginFormPasswordErrorBox">
              <p className = "mainLoginFormError">{passwordErrorMsg}</p>
            </div>
          </div>

          {/* 로그인 버튼 */}
          <button type = "submit" className = "mainLoginFormBtn">
              로그인
          </button>
        </form>
      </div>

    </div>
  )
}