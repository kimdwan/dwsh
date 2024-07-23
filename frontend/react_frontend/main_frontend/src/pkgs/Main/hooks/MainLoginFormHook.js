import * as yup from "yup"
import { yupResolver } from "@hookform/resolvers/yup"
import { useForm } from "react-hook-form"
import { MainLoginFormFetch } from "../functions"
import { useEffect, useState } from "react"

export const useMainLoginFormHook = () => {
  const [ emailErrorMsg, setEmailErrorMsg ] = useState(null)
  const [ passwordErrorMsg, setPasswordErrorMsg ] = useState(null)

  const schema = yup.object({
    email : yup.string().email("이메일 형식을 지켜주세요").required("이메일은 필수로 입력해야 합니다."),
    password : yup.string().min(4, "비밀번호는 최소 4글자 입니다.").max(16, "비밀번호는 최대 16글자 입니다.").required("비밀번호는 필수로 입력해야 합니다."),
  })

  const { register, handleSubmit, formState : {errors} } = useForm({
    resolver : yupResolver(schema)
  })

  const onSubmit = async (datas) => {
    const go_backend_url = process.env.REACT_APP_GO_BACKEND_URL
    const url = `${go_backend_url}/user/login`
    const response = await MainLoginFormFetch(url, datas, setEmailErrorMsg, setPasswordErrorMsg)
    if (response) {
      alert(response["message"])
    }
  }

  useEffect(() => {
    if (errors.email?.message) {
      if (!errors.password?.message) {
        setPasswordErrorMsg(null)
      }
      setEmailErrorMsg(errors.email.message)
    } else if (errors.password?.message) {
      if (!errors.email?.message) {
        setEmailErrorMsg(null)
      }
      setPasswordErrorMsg(errors.password.message)
    }

  }, [ errors ])


  return { register, handleSubmit, onSubmit, emailErrorMsg, passwordErrorMsg }
}