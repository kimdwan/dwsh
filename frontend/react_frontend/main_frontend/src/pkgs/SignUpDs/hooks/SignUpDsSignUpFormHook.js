import * as yup from "yup"
import { yupResolver } from "@hookform/resolvers/yup"
import { useForm } from "react-hook-form"
import { SignUpDsSignUpFormFetch } from "../functions"
import { useNavigate } from "react-router-dom"

export const useSignUpDsSignUpFormHook = (userAgree3, userType) => {
  const navigate = useNavigate()

  const schema = yup.object({
    email : yup.string().email("이메일 형식을 지켜주시길 바랍니다.").required("이메일은 필수적으로 입력해야 하는 사항입니다."),
    password : yup.string().min(4, "비밀번호는 최소 4글자 입니다.").max(16, "비밀번호는 최대 16글자 입니다.").required("비밀번호는 필수적으로 입력해야 하는 사항입니다."),
    confirm_password : yup.string().oneOf([yup.ref("password")],"비밀번호를 다시 확인해주십시오."),
    user_name : yup.string().max(10, "이름은 최대 10글자 입니다.").required("이름은 필수적으로 입력해야 하는 사항입니다."),
    our_first_day : yup.date().transform((value, originValue) => { return originValue === "" ? null : value }).required("만난날은 필수적으로 입력해야 하는 사항입니다."),
    secret_key : yup.string().max(30,"비밀키는 30글자를 넘을수 없습니다.").required("비밀키는 필수적으로 입력해야 하는 사항입니다."),
  })

  const { register, handleSubmit, formState : {errors}, setError } = useForm({
    resolver : yupResolver(schema)
  })

  const go_backend_url = process.env.REACT_APP_GO_BACKEND_URL

  const onSubmit = async (data) => {

    // 확인용 비밀번호 삭제
    delete data["confirm_password"]

    // 동의 사항 체크
    Array.from(["term_agree_1" ,"term_agree_2", "term_agree_3"]).forEach(( term_name, idx ) => {
      if (idx !== 2) {
        data[term_name] = true
      } else {
        data[term_name] = userAgree3
      }
    })

    // 유저의 타입을 넣는 함수
    data["user_type"] = userType

    // 시간 부분은 제거하는 함수 
    if (data["our_first_day"]) {
      const our_first_date = new Date(data["our_first_day"])
      const our_first_day = our_first_date.toISOString().split("T")[0]
      data["our_first_day"] = our_first_day
    }
    
    const url = `${go_backend_url}/user/signup`

    const response = await SignUpDsSignUpFormFetch(url, data, setError)
    if (response) {
      alert(response["message"])
      navigate("/")
    }
    
  }

  return { register, handleSubmit, errors, onSubmit }
}