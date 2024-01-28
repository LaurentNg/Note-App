import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Clipboard } from '@angular/cdk/clipboard';


@Component({
  selector: 'app-decrypt-text',
  templateUrl: './decrypt-text.component.html',
  styleUrl: './decrypt-text.component.scss'
})
export class DecryptTextComponent implements OnInit {
  notesToDecrypt = new FormGroup({
    text: new FormControl<string>('', [Validators.required]),
    passphrase: new FormControl<string>('', [Validators.required])
  });
  decryptedText = ""
  constructor(private clipboard: Clipboard) { }

  ngOnInit(): void {
  }

  onSubmit(): void {
    if (this.notesToDecrypt.valid) {
      const text = this.notesToDecrypt.get('text')?.value
      const passphrase = this.notesToDecrypt.get('passphrase')?.value
      console.log(text)
      console.log(passphrase)
      this.decryptedText = text + "hey"
    }
    else console.log('not valid')
  }

  copyDecryptedText() {
    const pending = this.clipboard.beginCopy(this.decryptedText);
    let remainingAttempts = 3;
    const attempt = () => {
      const result = pending.copy();
      if (!result && --remainingAttempts) {
        setTimeout(attempt);
      } else {
        pending.destroy();
      }
    };
    attempt();
  }
}
