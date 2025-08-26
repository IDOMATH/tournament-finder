import { Component, inject } from "@angular/core";
import { FormControl, FormGroup, Validators } from "@angular/forms";
import { valuesMatch } from "../../validators/value-match";
import { UserService } from "../../services/user-service";

@Component({
  selector: "app-sign-up",
  imports: [],
  templateUrl: "./sign-up.component.html",
  styleUrl: "./sign-up.component.css",
})
export class SignUpComponent {
  private userService = inject(UserService);
  form = new FormGroup({
    email: new FormControl("", { validators: [Validators.required] }),
    passwords: new FormGroup(
      {
        password: new FormControl("", { validators: [Validators.required] }),
        cPassword: new FormControl("", { validators: [Validators.required] }),
      },
      { validators: [valuesMatch("password", "cPassword")] }
    ),
  });

  onSubmit() {
    this.userService.signUp(
      this.form.controls.email.value!,
      this.form.controls.passwords.controls.password.value!
    );
  }
}
