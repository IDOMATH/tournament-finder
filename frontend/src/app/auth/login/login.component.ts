import { Component } from "@angular/core";
import { FormControl, FormGroup, Validators } from "@angular/forms";

@Component({
  selector: "app-login",
  imports: [],
  templateUrl: "./login.component.html",
  styleUrl: "./login.component.css",
})
export class LoginComponent {
  form = new FormGroup({
    email: new FormControl({ validators: [Validators.required] }),
    password: new FormControl({ validators: [Validators.required] }),
  });

  onSubmit() {
    console.log("Submitted");
  }
}
