import { useForm } from "react-hook-form";

type PatientFormValues = {
  full_name: string;
  date_of_birth: string;
  gender: "male" | "female" | "other";
  phone: string;
  email: string;
  address: string;
};

type Props = {
  onSubmit: (values: PatientFormValues) => Promise<void>;
};

export function PatientForm({ onSubmit }: Props) {
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting }
  } = useForm<PatientFormValues>({
    defaultValues: {
      gender: "female"
    }
  });

  return (
    <form className="grid gap-4 md:grid-cols-2" onSubmit={handleSubmit(onSubmit)}>
      <label className="grid gap-2">
        <span className="text-sm font-semibold text-slate-600">Full name</span>
        <input className="rounded-2xl border border-slate-200 px-4 py-3" {...register("full_name", { required: true })} />
        {errors.full_name ? <span className="text-sm text-coral">Name is required.</span> : null}
      </label>

      <label className="grid gap-2">
        <span className="text-sm font-semibold text-slate-600">Date of birth</span>
        <input type="date" className="rounded-2xl border border-slate-200 px-4 py-3" {...register("date_of_birth", { required: true })} />
      </label>

      <label className="grid gap-2">
        <span className="text-sm font-semibold text-slate-600">Gender</span>
        <select className="rounded-2xl border border-slate-200 px-4 py-3" {...register("gender", { required: true })}>
          <option value="female">Female</option>
          <option value="male">Male</option>
          <option value="other">Other</option>
        </select>
      </label>

      <label className="grid gap-2">
        <span className="text-sm font-semibold text-slate-600">Phone</span>
        <input className="rounded-2xl border border-slate-200 px-4 py-3" {...register("phone", { required: true })} />
      </label>

      <label className="grid gap-2">
        <span className="text-sm font-semibold text-slate-600">Email</span>
        <input type="email" className="rounded-2xl border border-slate-200 px-4 py-3" {...register("email")} />
      </label>

      <label className="grid gap-2 md:col-span-2">
        <span className="text-sm font-semibold text-slate-600">Address</span>
        <textarea className="min-h-24 rounded-2xl border border-slate-200 px-4 py-3" {...register("address", { required: true })} />
      </label>

      <div className="md:col-span-2">
        <button type="submit" disabled={isSubmitting} className="rounded-full bg-ink px-5 py-3 font-semibold text-white transition hover:bg-slate-800">
          {isSubmitting ? "Saving..." : "Create patient"}
        </button>
      </div>
    </form>
  );
}
