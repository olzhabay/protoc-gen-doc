// AUTOGENERATED CODE. DO NOT EDIT.

package gendoc

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
)

var embeddedResources = map[string]string{
	"docbook.tmpl": "H4sIAAAAAAAA/+xZ3W/bNhB/z19BaC/bikodugDDQLvA7LnF0GZB3O2dls42MYrUSMqNoel/Hyiq+v5a7bhZkJcguvvxjryP39ESfnMfMnQAqajgM+cH95WDgPsioHw3c/74uHr5k/NmfoWJ1NRnML9CCGuqGcxvpdDCFwwthR+HwDXRVHDsWe0VQkkiCd8BcleUgUpTs1SBb1BGXRhKEveGhJCmlbVmdUQkQe4SlC9pZFZlJip2P4BSZJebLo0jGsycJHFXMWPWsGNNVj2+F3zX4XXIr9HRLXLfEbWiwAJVyLEmGwZoK0kIM4cwVjgsXGKfEaU4CVveSwWyZhsbMiZ2UsQR8gVTM+fHinGEsBFG4BvlJxro/cz53vFORrxyr8dBr5sQvQcSVCUIYSk+1SUIYeBaHufZabFnH7ohH48RDCPekw2wYUglk51A7DX2iL3WQbDeiKCxrlLftWoYPXil4If2jRnlfyHzB3hZ0SYkpqLzKrKP2DOw+bA9s8JEa8xvVuXfUh7APXJ/z+KmkBNAJMEnGgLnnwC2JGYabQlT8F2aYgijPVFUzZcFysVeIU0S4EGadraW7amltfgnYbE5jsHNc9nPKEmaei8D5GanJdXkK4PXE11PK/ZsqxV84WWdXXJD1ULBBr/ea+CGP8/PCDegNASo9DBCDteXIIfHQR9FTE6lkF+IGkHcxOEG5FdmmY4qG43RV2Katr2F4JpQTvmuYblU/Hc2s2kZp7MnxDrYq92gqqqyUHgcXu5OdC6my4I8Rm+vL0FvJ/KSOdv/gE5svM9OJae15aWaqqeN8oe+sV7/9VKUt/l585LBAVj/nMaUb4UMCRvsludJ/jzJnyf5k5rktb6fQj55jaxBHqj/Ze82Hn6Gd8zvD6D34nG8vHhwwrJnReOT/g7+jkFpNE5dd6AiwRVMgD44P+WpvAA55fFpMEkubdGIbfVcu9YSSEj5Lk2Ryv4faufBPdjItzZhxb27sOov3MZjvP00FBVe6nxtu/YJI9Le27Oqrff++I2n/74z1tMd+usxwHn17YZqZSnPtBtJoUV/W3++mAhtAtgPWLx4MaT+jRzIkP72qPd9fGFlb8Wg+28Grb+7HVLfxZtjh75R2i2aapNUOR6z4qtTVF8KPs/M7INFpcMrz0ObN+RmkjOKWkTRJGsmVZOANmeToG+nHWSx3hMZjTveTzuJyWsvsEVcTdpqkFadsjpuThV2usJe8UHq3wAAAP//9p3zJMIaAAA=",
	"html.tmpl": "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",
	"markdown.tmpl": "H4sIAAAAAAAA/+RWTW/jNhC961dMpR42Xsi5L2wfGje7KHbTYB30EhQNbY1tATKpipSRQOR/L/ghkZZkJy3S0/pgiTPifLz3hlIC9xUTbMMKWLJNfUAqiMgZjWYEKDngPBasjBeza7KIoiSBB7IuENgWbhgVSAWPmqYidIcwvc0L5EpFTfPzNi/wL70dPs1hekcOqFQKj03j7v/8kLh7kEDoZs8qpa4igKZJId/C9BtyTnbIQSljdSlas1IANt5XRndhzNu6KEbjIs1crBSQZpB2K53vV1of+smM7f0zPQukPGd0kK5zuJwaz7TAIxbgfSZ3gK9PnmL30FvqWGF1zDcDhFuz7/s/99xaU3hcbUhBKviDFDXCw0uJug9ujOlRG1OhjVfRm9Xk9TkUkpPrrARS5Ds6j6t8txfxYkZgX+F2HidG0w+s1M/Nrksr7a7XqGmmS+SbKi/1KCgVlOUVeFLBCDB+aExoL5/R8JqSL4Tf5lhkOrgEcwvSwAUSvpI1FiAh2AkykpDqH9grnC7dD2SIqo7vpCB9yyA7jet8Id8uf8e33WbKMftM7R9ymuEzTH83hXGIMywr3BCBWSwz3JK6ELAlBccrpSaTZeedTiatfJqGsnUFp9hYZJY2gtGPUuCWn8AgGbpcKNOyVV979RiHgxZJP1se618I15e7+rDG6hzmQ9w72M/jH0z/GAcnFNi1PmdJTnO663tseW71P0NnXbOfUjveh/Z4TtNFMBzuxHzPyZDgNl8k4w3Imy7Pov4algEqQzTQvDwcFGMiC/B47egOsRl9Afy4in0aSPbprGY9OSdc9PTqX3b/UrIX5PoNxZ5lrWq/4981ctHS9B15ySjHdn2Wpj4j/WV/LcM3lC7g/BnvSuof9c48duJbWtwDK1EhOeR0pxRwc9+R4MLbFofxrf1SAvvEpQyvDeeAft5+5jjm3QwnCQy/STR901J/Drf83DGBHCTcfPwIEn4jRwIS7l/E3szeZ6ZdiTZ9udfs1uuXczTa6+jY+T/vDzRqyvR0hkI1n+49MGO4XkDo64ZP99LNaFmGPt1ZuLYthpbPJ7FuVntSld3T+5NgGoZ23SH+TwAAAP//htfnl2gMAAA=",
	"scalars.json": "H4sIAAAAAAAA/9yXzW4aMRDH7zzFiFMqBZDSlEa9JZGQOOQEOUWp5GVnvW6NTewxzaqq1HfoG/ZJqt0F1gYvpChVk9zQfHg9v/nPWNx1AL53AAC6C6NJT4sFdj9BN9Uukdg9rV1KE9rSvDbMFot45MzG7XxzciY1o+H52vGFLVk8ZZEvgpyNvaBcq6jLuKRYO0aVowPw47SlxiB1X4lBYFNhYN4q8P1ZrMCwjn9dn1DeNTb13Vq0sGRGsERiT6LilAOqmU6F4n0YK8wyMROoCDJtNh5QyBmJJYJy8wSNhd8/f4HIoNDOQCZQpiAsSPEVZQGkIWdl7DppyaRDewrOItjqYiCUJWRpPwI8uHkDXKgI7iDWh+1Fe6iFIuRo4rC9FB/1leDKzUEbGInH8tcJs2DwwQmD6btDPWi0/sJ6MDw/0IPm5k0PpFY82oT4SPvhu10YWDLCC9huxsBP3+3IXvLuGPlHOLgWMbqoGt0zy/EAgaM06Y4RZQuZmERcXCPu/4jkKET2qN05EVxhCkJRPWt9mOZoEebaIGxGWhZ1Cu6OM+VMgUHuJDNQXcG+7fVoj9qPz855eN7G+VWvwEw8YhqR8aX8xgoLWflqJAWh7cNNgK4G5FbvdLZ6OYAZBJ0RKuAGGaGp484+n128vM35zEqtWEakumKJgud0CGb54B6G+WH4epft/mF/uh7f9tL7Cy29wa2UaC2f8q/Lj2vK9K08eqZfZWlHpiKFbnuCCrecfoFT4/BaMmsHIyZt/XN/twOSTa+hdsDcWQJWd36mFTGh4HY66l2sXq+01NjHXiIILifX4zEQPlJMF+GHGmKhnbdczKc2CZuft3wiZGbJDJwS5ZVj3Ooz4aQqbf98VMrfAXbDig0fpgpgJhFkmCnA4oNDNSvXafvUtNG5KggnLYTu7svjYoR2s55OaR+dqsO9i6vxtEbUue/8CQAA//8z+wC/ohEAAA==",
}

func fetchResource(name string) ([]byte, error) {
	raw, ok := embeddedResources[name]
	if !ok {
		return nil, fmt.Errorf("Could not find resource for '%s'", name)
	}

	compressed, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	buf := bytes.NewBuffer(compressed)
	
	r, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(&out, r); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}
