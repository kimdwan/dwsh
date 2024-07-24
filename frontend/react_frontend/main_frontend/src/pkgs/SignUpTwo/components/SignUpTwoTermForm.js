import { useSignUpTwoClickTermBtnHook } from "../hooks"


export const SignUpTwoTermForm = ({ setTerms }) => {
  useSignUpTwoClickTermBtnHook(setTerms)

  return (
    <div className = "SignUpTwoTermFormContainer">
      
      {/* 이용 약간을 체크하는 장소 */}
      {
        Array.from(["one", "two", "three"]).map((count, idx) => {
          return (
            <div key = {idx} className = {`SignUpTwoTermFormBox${count}`}>
              
              {/* 머리에 해당함 */}
              <div className = {`SignUpTwoTermFormHeaderBox${count}`}>
                <input 
                  type = "checkbox"
                  id = {`SignUpTwoTermFormHeaderInput${count}`}
                  className = "SignUpTwoTermFormHeaderInput"
                />
                <div className = {`SignUpTwoTermFormHeaderText${count}`}>
                  {
                    idx < 2 ? `필수${idx + 1}` : `선택`
                  }
                </div>
              </div>

              {/* 안에들어가는 동의 내용 */}
              <div className = {`SignUpTwoTermFormBodyBox${count}`}>
                <textarea className = {`SignUpTwoTermFormBodyText${count}`} value = {
                  idx === 0 ? `
                  저희 홈페이지에 오신걸 환영합니다 해당 사항은 필수로 동의 하셔야 하는 부분입니다.
                  ` : idx === 1 ? `
                    지금 부터 저희 홈페이지에서 지켜야 할 사항들을 말씀드리겠습니다.
                  ` : `
                    이건 선택 사항으로 체크하셔도 되고 안하셔도 됩니다.
                  `
                } readOnly/>
              </div>

            </div>
          )
        })
      }

    </div>
  )
}