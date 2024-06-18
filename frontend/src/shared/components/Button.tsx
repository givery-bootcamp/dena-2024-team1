import classNames from "classnames";
import type { ButtonHTMLAttributes, FC } from "react";

type Variant = "primary" | "secondary" | "alert";

type Props = {
  variant?: Variant
} & ButtonHTMLAttributes<HTMLButtonElement>

const variantToBgColor: Record<Variant, string> = {
  primary: "bg-primary",
  secondary: "bg-gray-100",
  alert: "bg-alert",
};

export const Button: FC<Props> = ({ variant = "primary", ...props }) => {
  const bgColorClassName = variantToBgColor[variant];
  const buttonClassName = "text-white font-semibold py-2 px-8 rounded";

  return <button {...props} className={classNames(bgColorClassName, buttonClassName, props.className)} />;
};
