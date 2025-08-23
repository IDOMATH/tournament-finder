import { Component, inject } from "@angular/core";
import { FormControl, FormGroup, Validators } from "@angular/forms";
import { UserService } from "../../services/user-service";

@Component({
  selector: "app-login",
  imports: [],
  templateUrl: "./login.component.html",
  styleUrl: "./login.component.css",
})
export class LoginComponent {
  private userService = inject(UserService);
  form = new FormGroup({
    email: new FormControl("", { validators: [Validators.required] }),
    password: new FormControl("", { validators: [Validators.required] }),
  });

  onSubmit() {
    this.userService.login(
      this.form.controls.email.value!,
      this.form.controls.password.value!
    );
  }
}
