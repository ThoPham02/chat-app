import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";

// Schema validation bằng Yup
const schema = yup.object().shape({
  fullName: yup
    .string()
    .min(2, "Họ và tên phải có ít nhất 2 ký tự")
    .required("Vui lòng nhập họ và tên"),
  phoneNumber: yup
    .string()
    .matches(/^0[0-9]{9}$/, "Số điện thoại không hợp lệ")
    .required("Vui lòng nhập số điện thoại"),
  password: yup
    .string()
    .min(6, "Mật khẩu phải có ít nhất 6 ký tự")
    .required("Vui lòng nhập mật khẩu"),
  confirmPassword: yup
    .string()
    .oneOf([yup.ref("password")], "Mật khẩu xác nhận không khớp")
    .required("Vui lòng xác nhận mật khẩu"),
  cccd: yup
    .string()
    .matches(/^\d{12}$/, "Số CCCD không hợp lệ, phải có 12 chữ số")
    .required("Vui lòng nhập số CCCD"),
  issueDate: yup
    .date()
    .max(new Date(), "Ngày cấp không được lớn hơn ngày hiện tại")
    .required("Vui lòng chọn ngày cấp"),
  issuePlace: yup
    .string()
    .min(3, "Nơi cấp phải có ít nhất 3 ký tự")
    .required("Vui lòng nhập nơi cấp"),
});

export default function RegisterPage() {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({
    resolver: yupResolver(schema),
  });

  const onSubmit = (data) => {
    console.log("Dữ liệu đăng ký:", data);
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-navy">
      <div className="bg-white p-8 rounded-lg shadow-lg border-2 border-orange w-full max-w-2xl">
        <h2 className="text-navy text-2xl font-bold text-center mb-6">
          Đăng Ký Tài Khoản
        </h2>

        <form onSubmit={handleSubmit(onSubmit)}>
          <div className="grid grid-cols-2 gap-4">
            {/* Họ và Tên */}
            <div>
              <label className="block text-navy font-medium">Họ và Tên</label>
              <input
                {...register("fullName")}
                className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-orange"
                placeholder="Nhập họ và tên"
              />
              <p className="text-red-500 text-sm">{errors.fullName?.message}</p>
            </div>

            {/* Số điện thoại */}
            <div>
              <label className="block text-navy font-medium">
                Số điện thoại
              </label>
              <input
                {...register("phoneNumber")}
                className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-orange"
                placeholder="Nhập số điện thoại"
              />
              <p className="text-red-500 text-sm">
                {errors.phoneNumber?.message}
              </p>
            </div>

            {/* Mật khẩu */}
            <div>
              <label className="block text-navy font-medium">Mật khẩu</label>
              <input
                type="password"
                {...register("password")}
                className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-orange"
                placeholder="Nhập mật khẩu"
              />
              <p className="text-red-500 text-sm">{errors.password?.message}</p>
            </div>

            {/* Xác nhận mật khẩu */}
            <div>
              <label className="block text-navy font-medium">
                Nhập lại mật khẩu
              </label>
              <input
                type="password"
                {...register("confirmPassword")}
                className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-orange"
                placeholder="Xác nhận mật khẩu"
              />
              <p className="text-red-500 text-sm">
                {errors.confirmPassword?.message}
              </p>
            </div>

            {/* Số CCCD */}
            <div>
              <label className="block text-navy font-medium">Số CCCD</label>
              <input
                {...register("cccd")}
                className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-orange"
                placeholder="Nhập số CCCD"
              />
              <p className="text-red-500 text-sm">{errors.cccd?.message}</p>
            </div>

            {/* Ngày cấp */}
            <div>
              <label className="block text-navy font-medium">Ngày cấp</label>
              <input
                type="date"
                {...register("issueDate")}
                className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-orange"
              />
              <p className="text-red-500 text-sm">
                {errors.issueDate?.message}
              </p>
            </div>

            {/* Nơi cấp */}
            <div className="col-span-2">
              <label className="block text-navy font-medium">Nơi cấp</label>
              <input
                {...register("issuePlace")}
                className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-orange"
                placeholder="Nhập nơi cấp"
              />
              <p className="text-red-500 text-sm">
                {errors.issuePlace?.message}
              </p>
            </div>
          </div>

          {/* Nút đăng ký */}
          <button
            type="submit"
            className="w-full bg-orange text-white py-2 rounded-lg hover:bg-orange-dark transition duration-300 mt-4"
          >
            Đăng Ký
          </button>

          {/* Đã có tài khoản? */}
          <div className="text-center mt-4">
            <a href="#" className="text-orange hover:underline">
              Đã có tài khoản? Đăng nhập
            </a>
          </div>
        </form>
      </div>
    </div>
  );
}
