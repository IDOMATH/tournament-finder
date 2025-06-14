import { Component } from "@angular/core";
import {
  AbstractControl,
  FormControl,
  FormGroup,
  Validators,
} from "@angular/forms";

function valuesMatch(controlName1: string, controlName2: string) {
  return (control: AbstractControl) => {
    const val1 = control.get(controlName1)?.value;
    const val2 = control.get(controlName2)?.value;

    if (val1 === val2) {
      return null;
    }

    return { passwordsDoNotMatch: true };
  };
}

@Component({
  selector: "app-sign-up",
  imports: [],
  templateUrl: "./sign-up.component.html",
  styleUrl: "./sign-up.component.css",
})
export class SignUpComponent {
  form = new FormGroup({
    email: new FormControl({ validators: [Validators.required] }),
    passwords: new FormGroup(
      {
        password: new FormControl({ validators: [Validators.required] }),
        cPassword: new FormControl({ validators: [Validators.required] }),
      },
      { validators: [valuesMatch("password", "cPassword")] }
    ),
  });

  onSubmit() {
    console.log("submitted");
  }
}
