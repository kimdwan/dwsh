import { useSignUpGuestGetUserTermAgree3AndUserTypeHook, useSignUpGuestSignUpFormHook } from "../hooks"

export const SignUpGuestSignUpForm = () => {
  const { termAgree3, userType } = useSignUpGuestGetUserTermAgree3AndUserTypeHook()
  const { register, handleSubmit, errors, onSubmit } = useSignUpGuestSignUpFormHook(termAgree3, userType)

  return (
    <div className = "SignUpGuestSignUpFormContainer">
      
      <div className = "SignUpGuestSignUpFormBox">
        <form onSubmit = {handleSubmit(onSubmit)}>

          {/* 본격적으로 회원가입과 관련된 작업이 진행되는 장소  */}
          <div className = "SignUpGuestSignUpFormDivBox">

            {/* 이메일 관련 */}
            <div className = "SignUpGuestSignUpFormEmailBox">
              {/* 실질적 장소 */}
              <div className = "SignUpGuestSignUpFormEmailInputBox">
                <div className = "SignUpGuestSignUpFormEmailText">
                  이메일
                </div>
                <input
                  {...register("email")}
                  type = "text"
                  className = "SignUpGuestSignUpFormEmailInputValue" 
                />
              </div>

              {/* 에러 처리 */}
              <div className = "SignUpGuestSignUpFormEmailErrorBox">
                {errors.email?.message && <p className = "SignUpGuestSignUpFormErrorMsg">{errors.email.message}</p>}
              </div>
            </div>

            {/* 비밀번호 관련 */}
            <div className = "SignUpGuestSignUpFormPasswordBox">
              {/* 실질적 장소 */}
              <div className = "SignUpGuestSignUpFormPasswordInputBox">
                <div className = "SignUpGuestSignUpFormPasswordText">
                  비밀번호
                </div>
                <input 
                  {...register("password")}
                  type = "password"
                  className = "SignUpGuestSignUpFormPasswordInputValue"
                />
              </div>
              {/* 에러 처리 */}
              <div className = "SignUpGuestSignUpFormPasswordErrorBox">
                { errors.password?.message && <p className = "SignUpGuestSignUpFormErrorMsg">{ errors.password.message }</p> }
              </div>
            </div>

            {/* 비밀번호 확인 관련 */}
            <div className = "SignUpGuestSignUpFormConfirmPasswordBox">
              {/* 실질적인 부분 */}
              <div className = "SignUpGuestSignUpFormConfirmPasswordInputBox">
                <div className = "SignUpGuestSignUpFormConfirmPasswordText">
                  비밀번호 확인
                </div>
                <input
                  {...register("confirm_password")}
                  className = "SignUpGuestSignUpFormConfirmPasswordInputValue"
                  type= "password" 
                />
              </div>
              {/* 에러 처리 */}
              <div className = "SignUpGuestSignUpFormConfirmPasswordErrorBox">
                { errors.confirm_password?.message && <p className = "SignUpGuestSignUpFormErrorMsg">{ errors.confirm_password.message }</p> }
              </div>
            </div>

            {/* 유저 이름 관련 */}
            <div className = "SignUpGuestSignUpFormUserNameBox">
              {/* 실질적인 부분 */}
              <div className = "SignUpGuestSignUpFormUserNameInputBox">
                <div className = "SignUpGuestSignUpFormUserNameText">
                  이름
                </div>
                <input
                  {...register("user_name")}
                  type = "text"
                  className = "SignUpGuestSignUpFormUserNameInputValue" 
                />
              </div>
              {/* 에러 처리 */}
              <div className = "SignUpGuestSignUpFormUserNameErrorBox">
                { errors.user_name?.message && <p className = "SignUpGuestSignUpFormErrorMsg">{ errors.user_name.message }</p> }
              </div>
            </div>
          </div>

          {/* 회원가입 버튼이 있는 장소 */}
          <div className = "SignUpGuestSignUpFormInputBox">
            <button className = "SignUpGuestSignUpFormInputBtn" type = "submit">
              회원가입
            </button>
          </div>

        </form>
      </div>

    </div>
  )
}