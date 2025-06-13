import { Component } from "@angular/core";
import { FormControl, FormGroup, Validators } from "@angular/forms";

@Component({
  selector: "app-sign-up",
  imports: [],
  templateUrl: "./sign-up.component.html",
  styleUrl: "./sign-up.component.css",
})
export class SignUpComponent {
  form = new FormGroup({
    email: new FormControl({ validators: [Validators.required] }),
    // TODO: Add password matching validator
    password: new FormControl({ validators: [Validators.required] }),
    cpassword: new FormControl({ validators: [Validators.required] }),
  });

  onSubmit() {
    console.log("submitted");
  }
}
