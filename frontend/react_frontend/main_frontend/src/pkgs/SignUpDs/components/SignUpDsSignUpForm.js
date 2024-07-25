import { useSignUpDsUserTerm3AndUserTypeHook, useSignUpDsSignUpFormHook } from "../hooks"

export const SignUpDsSignUpForm = () => {
  const { userAgree3, userType } = useSignUpDsUserTerm3AndUserTypeHook()
  const { register, handleSubmit, errors, onSubmit } = useSignUpDsSignUpFormHook(userAgree3, userType)
  
  return (
    <div className = "SignUpDsSignUpFormContainer">
      
      {/* 회원가입의 폼을 작성하는 장소 */}
      <div className = "SignUpDsSignUpFormBox">
        <form onSubmit = {handleSubmit(onSubmit)}>
          <div className = "SignUpDsSignUpFormDivBox">
            {/* 이메일 관련 폼 */}
            <div className = "SignUpDsSignUpFormEmailBox">
              {/* 실질적 입력 폼 */}
              <div className = "SignUpDsSignUpFormEmailInputBox">
                <div className = "SignUpDsSignUpFormEmailInputText">
                  이메일
                </div>
                <input 
                  {...register("email")}
                  type = "text"
                  className = "SignUpDsSignUpFormEmailInputValue"
                />
              </div>
              {/* 에러 처리 */}
              <div className = "SignUpDsSignUpFormEmailErrorBox">
                { errors.email?.message && <p className = "SignUpDsSignUpFormErrorMsg">{errors.email.message}</p> }
              </div>
            </div>

            {/* 비밀번호 관련 폼 */}
            <div className = "SignUpDsSignUpFormPasswordBox">
              {/* 실질적 입력 폼 */}
              <div className = "SignUpDsSignUpFormPasswordInputBox">
                <div className = "SignUpDsSignUpFormPasswordInputText">
                  비밀번호
                </div>
                <input
                  {...register("password")}
                  type = "password"
                  className = "SignUpDsSignUpFormPasswordInputValue"
                />
              </div>
              {/* 에러 처리 */}
              <div className = "SignUpDsSignUpFormPasswordErrorBox">
                { errors.password?.message && <p className = "SignUpDsSignUpFormErrorMsg">{errors.password.message}</p> }
              </div>
            </div>

            {/* 비밀번호 확인 관련폼 */}
            <div className = "SignUpDsSignUpFormConfirmPasswordBox">
              {/* 실질적 입력 폼 */}
              <div className = "SignUpDsSignUpFormConfirmPasswordInputBox">
                <div className = "SignUpDsSignUpFormConfirmPasswordInputText">
                  비밀번호 확인
                </div>
                <input 
                  {...register("confirm_password")}
                  type = "password"
                  className = "SignUpDsSignUpFormConfirmPasswordInputValue"
                />
              </div>
              {/* 에러 처리 */}
              <div className = "SignUpDsSignUpFormConfirmPasswordErrorBox">
                { errors.confirm_password?.message && <p className = "SignUpDsSignUpFormErrorMsg">{errors.confirm_password.message}</p> }
              </div>
            </div>

            {/* 유저의 이름 관련폼 */}
            <div className = "SignUpDsSignUpFormUserNameBox">
              {/* 실질적 입력 폼 */}
              <div className = "SignUpDsSignUpFormUserNameInputBox">
                <div className = "SignUpDsSignUpFormUserNameInputText">
                  이름
                </div>
                <input
                  {...register("user_name")}
                  className = "SignUpDsSignUpFormUserNameInputValue"
                  type = "text" 
                />
              </div>
              {/* 에러 처리 */}
              <div className = "SignUpDsSignUpFormUserNameErrorBox">
                { errors.user_name?.message && <p className = "SignUpDsSignUpFormErrorMsg">{ errors.user_name.message }</p> }
              </div>
            </div>  

            {/* 만난 날 관련 폼 */}
            <div className = "SignUpDsSignUpFormOurFirstDayBox">
              {/* 실질적 입력 폼 */}
              <div className = "SignUpDsSignUpFormOurFirstDayInputBox">
                <div className = "SignUpDsSignUpFormOurFirstDayInputText">
                  만난날
                </div>
                <input 
                  {...register("our_first_day")}
                  className = "SignUpDsSignUpFormOurFirstDayInputValue"
                  type = "date"
                />
              </div>
              {/* 에러 처리 */}
              <div className = "SignUpDsSignUpFormOurFirstDayErrorBox">
                { errors.our_first_day?.message && <p className = "SignUpDsSignUpFormErrorMsg">{errors.our_first_day.message}</p> }
              </div>
            </div>

            {/* 비밀키 입력 폼 */}
            <div className = "SignUpDsSignUpFormSecretKeyBox">
              {/* 실질적 입력 폼 */}
              <div className = "SignUpDsSignUpFormSecretKeyInputBox">
                <div className = "SignUpDsSignUpFormSecretKeyInputText">
                  비밀키
                </div>
                <input 
                  {...register("secret_key")}
                  className = "SignUpDsSignUpFormSecretKeyInputValue"
                  type = "password"
                />
              </div>
              {/* 에러 처리 */}
              <div className = "SignUpDsSignUpFormSecretKeyErrorBox">
                { errors.secret_key?.message && <p className = "SignUpDsSignUpFormErrorMsg">{errors.secret_key.message}</p> }
              </div>
            </div>
          </div>

          {/* 회원가입을 처리하는 박스 */}
          <div className = "SignUpDsSignUpFormInputBox">
            <button className = "SignUpDsSignUpFormInputBtn" type = "submit">
              회원가입
            </button>
          </div>

        </form>
      </div>

    </div>
  )
}